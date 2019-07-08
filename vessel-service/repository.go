package main

import (
	"context"
	vesselProto "github.com/Incognida/shippy_protos/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(specification *vesselProto.Specification) (*vesselProto.Vessel, error)
	Create(vessel *vesselProto.Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (r *MongoRepository) FindAvailable(spec *vesselProto.Specification) (*vesselProto.Vessel, error) {
	cur, err := r.collection.Find(context.Background(), nil, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var vessel *vesselProto.Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		} else if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, nil
}

func (r *MongoRepository) Create(vessel *vesselProto.Vessel) error {
	_, err := r.collection.InsertOne(context.Background(), vessel)
	return err
}
