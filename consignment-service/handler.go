// shippy-service-consignment/handler.go
package main

import (
	"context"
	pb "github.com/Incognida/shippy_protos/consignment"
	vesselProto "github.com/Incognida/shippy_protos/vessel"
	"log"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	v, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}

	if v.Vessel != nil {
		log.Printf("Found vessel: %v \n", v)
		req.VesselId = v.Vessel.Id
	}

	// Save our consignment
	if err = s.repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
