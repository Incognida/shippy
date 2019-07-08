// shippy-service-consignment/main.go
package main

import (
	"context"
	"fmt"
	pb "github.com/Incognida/shippy_protos/consignment"
	vesselProto "github.com/Incognida/shippy_protos/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vs := grpc.NewService()
	vs.Init()

	vesselClient := vesselProto.NewVesselService("vessel", vs.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
