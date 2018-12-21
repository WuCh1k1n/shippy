build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
	docker build -t shippy/consignment-cli .

run:
	docker run --rm -e MICRO_REGISTRY=mdns \
	    --name consignment-cli shippy/consignment-cli