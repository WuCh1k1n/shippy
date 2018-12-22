package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
	vesselPb "com.fengberlin/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Close()

	// 创建一个新服务，其中可以包括一些可选的配置
	srv := micro.NewService(
		// 这个 Name 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init 会解析命令行参数
	srv.Init()

	// 注册 handler，相当于是之前的 gRPC server
	// 第二个参数需要传入实现了 ShippingServiceHandler 接口的对象
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}