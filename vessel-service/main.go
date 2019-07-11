// vessel-service/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/Incognida/shippy_protos/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	srv := micro.NewService(micro.Name("vessel"))
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("shippy").Collection("vessels")

	r := &MongoRepository{collection: collection}
	err = r.Create(&pb.Vessel{
		Capacity:  69,
		MaxWeight: 69000,
		Name:      "Chaika",
		Available: true,
		OwnerId:   "Pistoletov",
	})

	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterVesselServiceHandler(srv.Server(), &handler{r})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
