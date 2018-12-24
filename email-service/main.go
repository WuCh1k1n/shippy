package main

import (
	"context"
	"log"

	userPb "com.fengberlin/shippy/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

type Subscriber struct{}

func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

func sendEmail(user *userPb.User) error {
	log.Printf("[SENDING A EMAIL TO %s...]", user.Name)
	return nil
}

func (s *Subscriber) Process(ctx context.Context, user *userPb.User) error {

	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]:", user.Name)
	return nil
}