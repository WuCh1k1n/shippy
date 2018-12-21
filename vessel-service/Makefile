build:
	protoc -I proto/vessel --go_out=plugins=micro:proto/vessel proto/vessel/vessel.proto
	# 由于国内的环境，所以这里不使用多阶段构建了，而选择直接build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	docker build -t shippy/vessel-service .

run:
	docker run --rm -d -p 50052:50051 \
	    -e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		--name vessel-service shippy/vessel-service