package main

import (
	"context"

	pb "com.fengberlin/shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	defer s.GetRepo().Close()

	vessel, err := s.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	defer s.GetRepo().Close()

	err := s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	resp.Vessel = req
	resp.Created = true
	return nil
}