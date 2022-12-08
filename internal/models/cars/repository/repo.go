package repository

import (
	"context"
	"dig_practice/internal/models/cars"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CarRepo interface {
	SaveCar(car cars.Car) error
}

type CarRepoConfig struct {
	DBHost  string
	DBPort  string
	DBName  string
	DBTable string
}

type carRepo struct {
	collection *mongo.Collection
}

func (r *carRepo) SaveCar(car cars.Car) error {
	fmt.Println("saving car", car)
	return nil
}

func (r *carRepo) GetCars() error {
	return nil
}

func NewCarRepo(cfg *CarRepoConfig) *carRepo {
	fmt.Println("my Config", cfg)
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBTable)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database(cfg.DBName).Collection(cfg.DBTable)

	return &carRepo{coll}
}
