build:
	protoc -I proto/user --go_out=plugins=micro:proto/user proto/user/user.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	docker build -t shippy/user-service .

run:
	docker run -d --rm -p 50053:50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		--name user-service shippy/user-service