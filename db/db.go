package db

import (
	"context"
	"github.com/bhupeshpandey/cars/model"
	"strings"
)

type DB struct {
	db model.DB
}

func New(config *model.DBConfig) (model.DB, error) {

	var db model.DB
	var err error
	if strings.EqualFold(config.Env, "Testing") {
		db = newMockDB()
	} else {
		db, err =  newMongoDB(config, context.Background())

		if err != nil {
			return nil, err
		}
	}
	return &DB {
		db,
	}, nil
}

func (db DB) GetProducts() ([]model.Product, error) {
	return db.db.GetProducts()
}

func (db DB) GetCarWithId(id string) (model.Product, error) {
	return db.db.GetCarWithId(id)
}

func (db DB) GetAccessoryWithId(id string) (model.Product, error) {
	return db.db.GetAccessoryWithId(id)
}

func (db DB) GetAccessories() ([]model.Product, error) {
	return db.db.GetAccessories()
}
func (db DB) GetCars() ([]model.Product, error) {
	return db.db.GetCars()
}
func (db DB) GetRecentItems() ([]model.Product, error) {
	return db.db.GetRecentItems()
}
func (db DB) SetSelectedItem(product model.Product) (bool, error) {
	return db.db.SetSelectedItem(product)
}

