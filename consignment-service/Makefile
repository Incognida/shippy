build:
	protoc  --plugin=grpc:. \
	--micro_out=. --go_out=. proto/consignment/consignment.proto

	docker build -t shippy-service-consignment .
run:
	docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns \
        shippy-service-consignment
