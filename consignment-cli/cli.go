// consignment-cli/cli.go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
)

const (
	defaultFilename = "consignment.json"
)

// parseFile: 解析json文件得到需要创建的consignment的信息
func parseFile(filename string) (*pb.Consignment, error) {

	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	cmd.Init()

	// 创建 client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	filename := defaultFilename
	if len(os.Args) < 3 {
		log.Fatalln("Not enough arguments, expecing file and token.")
	}
	filename = os.Args[1]

	token := os.Args[2]

	consignment, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Could not parse file: %v\n", err)
	}

	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v\n", err)
	}
	// 是否成功创建consignment
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v\n", err)
	}
	// 查看所有的consignments
	for _, consignment := range getAll.Consignments {
		log.Println(consignment)
	}
}