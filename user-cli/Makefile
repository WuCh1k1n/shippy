build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	docker build -t shippy/user-cli .

run:
	docker run --rm -e MICRO_REGISTRY=mdns \
	    --name user-cli shippy/user-cli