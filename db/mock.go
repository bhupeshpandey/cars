package db

import "github.com/bhupeshpandey/cars/model"

type mockDB struct {
}

func newMockDB() *mockDB {
	return &mockDB {}
}

func (mockdb *mockDB) GetProducts() ([]model.Product, error) {
	return nil, nil
}

func (mockdb *mockDB) GetCarWithId(id string) (model.Product, error) {
	return model.Product{}, nil
}

func (mockdb *mockDB) GetAccessoryWithId(id string) (model.Product, error) {
	return model.Product{}, nil
}
func (mockdb *mockDB) GetAccessories() ([]model.Product, error) {
	return nil, nil
}
func (mockdb *mockDB) GetCars() ([]model.Product, error) {
	return nil, nil
}
func (mockdb *mockDB) GetRecentItems() ([]model.Product, error) {
	return nil, nil
}
func (mockdb *mockDB) SetSelectedItem(product model.Product) (bool, error) {
	return false, nil
}
