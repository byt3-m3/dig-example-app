package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type ModelWriter interface {
	SaveModel(ctx context.Context, model interface{}) (isSuccess bool, err error)
}

type modelWriter struct {
	collection MongoCollection
}

func (w *modelWriter) SaveModel(ctx context.Context, model interface{}) (isSuccess bool, err error) {
	result, err := w.collection.InsertOne(ctx, model)
	if err != nil {
		log.Println("error inserting document")
		return false, err

	}
	resultObjID := result.InsertedID.(primitive.ObjectID)
	if resultObjID == primitive.NilObjectID {
		return false, nil
	}
	return true, nil
}
