package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/micro/go-micro/broker"

	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"

	pb "com.fengberlin/shippy/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
	pubSub       broker.Broker
}

func (srv *service) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {

	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}

	resp.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {

	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {

	log.Println("Logging in with:", req.Email)
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	log.Println(user)

	// 将我们给定的密码与存储在数据库中的散列密码进行比较
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	resp.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {

	// 生成密码的散列版本
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPwd)

	if err := srv.repo.Create(req); err != nil {
		return err
	}

	resp.User = req

	if err := srv.publishEvent(req); err != nil {
		return err
	}

	return nil
}

const topic = "user.created"

// 发布消息
func (srv *service) publishEvent(user *pb.User) error {

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	// 发布 user.created topic 消息
	if err := srv.pubSub.Publish(topic, msg); err != nil {
		log.Fatalf("[pub] failed: %v\n", err)
	}

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {

	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}
