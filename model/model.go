package model

type AppConfig struct {
	Server *ServerConfig
	Cache *CacheConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type CacheConfig struct {
	Size int
}

type Cache interface {
	Get(id string) string
	Put(id, name string)
}

type Item struct {
	Id string
	Name string
}
