package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
)

// IRepository - 存储库接口
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - 虚拟存储库，模拟数据存储的使用，以后会用一个真正的实现来替换
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service 实现所有方法以满足我们在protobuf定义中定义的服务。
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 可以去 consignment.pb.go 中寻找对应方法的签名以更好地实现它们
type service struct {
	repo IRepository
}

// CreateConsignment - 我们 rpc 服务中的一个方法
// 实现 ShippingServiceServer 接口
// 使用grpc生成出来的函数签名：func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error)

// 下面是使用go micro插件生成出来的新代码，这消除了样板式代码
// 这次实现的是 ShippingServiceHandler 接口
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.repo.GetAll()
	resp.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	// 创建一个新服务，其中可以包括一些可选的配置
	srv := micro.NewService(
		// 这个 Name 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init 会解析命令行参数
	srv.Init()

	// 注册 handler，相当于是之前的 gRPC server
	// 第二个参数需要传入实现了 ShippingServiceHandler 接口的对象
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}