// consignment-cli/cli.go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "com.fengberlin/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
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
	// 建立与gPRC server的连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	defer conn.Close()

	// 建立gRPC client
	client := pb.NewShippingServiceClient(conn)

	filename := defaultFilename
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	consignment, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Could not parse file: %v\n", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v\n", err)
	}
	// 是否成功创建consignment
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v\n", err)
	}
	// 查看所有的consignments
	for _, consignment := range getAll.Consignments {
		log.Println(consignment)
	}
}