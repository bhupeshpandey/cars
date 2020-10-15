package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppConfig struct {
	ServerConfig *ServerConfig `json:'serverConfig'`
	CacheConfig  *CacheConfig `json:"cacheConfig"`
	DBConfig     *DBConfig `json:"dbConfig"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
}

type DBConfig struct {
	Env string `json:"env"`// defines the env for DB (Testing, Development, Production)
	Host string `json:"host"`
	Port int `json:"port"`
	// TODO Add user name and password going forward.
}

type CacheConfig struct {
	Size int `json:"capacity"`
}

type Cache interface {
	Get(id string) string
	Put(id, name string)
}

type DB interface {
	GetProducts() ([]Product, error)
	GetCarWithId(id string) (Product, error)
	GetAccessoryWithId(id string) (Product, error)
	GetAccessories()([]Product, error)
	GetCars()([]Product, error)
	GetRecentItems()([]Product, error)
	SetSelectedItem(product Product) (bool, error)
}

type Product struct {
	ID    primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name,omitempty" bson:"name,omitempty"`
	Price float64 `json:"price" bson:"price,omitempty`
	Type  string  `json:"type" bson:"type,omitempty"`
}
