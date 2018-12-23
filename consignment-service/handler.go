// gRPC 服务需要实现的接口逻辑
package main

import (
	"context"
	"log"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
	vesselPb "com.fengberlin/shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	// clone 主会话
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// 关闭 mgo session 连接
	defer s.GetRepo().Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s\n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id

	err = s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	defer s.GetRepo().Close()

	consignments, err := s.GetRepo().GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}