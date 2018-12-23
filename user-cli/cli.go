package main

import (
	"context"
	"log"
	"os"

	pb "com.fengberlin/shippy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

func main() {

	cmd.Init()

	// 创建 user-service 客户端
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// 暂时将用户信息写死在代码中
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	resp, err := client.Create(context.Background(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("call Create method error: %v\n", err)
	}
	log.Println("user created: ", resp.User.Id)

	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("call GetAll method error: %v", err)
	}
	for i, u := range allResp.Users {
		log.Printf("user_%d: %v\n", i, u)
	}

	authResp, err := client.Auth(context.Background(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v\n", err)
	}
	log.Println("token: ", authResp.Token)

	os.Exit(0)
}