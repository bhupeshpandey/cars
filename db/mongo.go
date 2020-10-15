package db

import (
	"context"
	"fmt"
	"github.com/bhupeshpandey/cars/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	host   string
	port   int
	ctx    context.Context
	client *mongo.Client
	carsCollection *mongo.Collection
	accessoriesCollection *mongo.Collection
}

func newMongoDB(db *model.DBConfig, ctx context.Context) (*mongoDB, error) {
	ctxx, _ := context.WithCancel(ctx)
	 mongoInst := &mongoDB{
	 	db.Host,
	 	db.Port,
		ctxx,
		nil,
		nil,
		nil,
	 }

	 if err := mongoInst.Init(); err != nil {
	 	return nil, err
	 }
	 return mongoInst, nil

}

func (mongoDB *mongoDB) Init() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%v/", mongoDB.host, mongoDB.port))
	client, err := mongo.Connect(mongoDB.ctx, clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(mongoDB.ctx, nil)
	if err != nil {
		return err
	}

	mongoDB.client = client

	database := client.Database("Products")

    mongoDB.carsCollection = database.Collection("Cars")
	mongoDB.accessoriesCollection = database.Collection("Accessories")
	return nil
}



func (mongoDB *mongoDB) GetProducts() ([]model.Product, error) {
	return nil, nil
}

func (mongoDB *mongoDB) GetCarWithId(id string) (model.Product, error) {
	return model.Product{}, nil
}

func (mongoDB *mongoDB) GetAccessoryWithId(id string) (model.Product, error) {
	return model.Product{}, nil
}
func (mongoDB *mongoDB) GetAccessories() ([]model.Product, error) {
	return nil, nil
}
func (mongoDB *mongoDB) GetCars() ([]model.Product, error) {
	return nil, nil
}
func (mongoDB *mongoDB) GetRecentItems() ([]model.Product, error) {
	return nil, nil
}
func (mongoDB *mongoDB) SetSelectedItem(product model.Product) (bool, error) {
	return false, nil
}