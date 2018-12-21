package main

import (
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"context"
	"log"
	"net"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
)

const (
	port = ":50051"
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
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	repo := &Repository{}

	// 建立一个gRPC server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	// 向 gPRC server 注册服务
	pb.RegisterShippingServiceServer(s, &service{repo})

	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}