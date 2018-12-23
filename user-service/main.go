package main

import (
	"log"

	micro "github.com/micro/go-micro"

	pb "com.fengberlin/shippy/user-service/proto/user"
)

func main() {

	db, err := CreateConnection()
	if err != nil {
		log.Fatalln("can not obtain the connect of database: ", err)
	}
	defer db.Close()

	// 自动将 user struct 迁移到数据库列/类型等。
	// 这将检查更新并在每次重新启动此服务时迁移它们。
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}