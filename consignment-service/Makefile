# 一定要注意 Makefile 中的缩进(要跟一个tab，而不是4个空格)，否则执行 make build 可能报错
build:
	# protoc -I proto/consignment --go_out=plugins=grpc:proto/consignment proto/consignment/consignment.proto
	protoc -I proto/consignment --go_out=plugins=micro:proto/consignment proto/consignment/consignment.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	docker build -t shippy/consignment-service .

run:
	docker run -d --rm -p 50051:50051 \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    -e MICRO_REGISTRY=mdns \
	    --name consignment-service shippy/consignment-service