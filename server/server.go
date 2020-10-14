package server

import (
	//mux "github.com/gorilla/mux"
	//lru "github.com/cars/cache"
	model "github.com/bhupeshpandey/cars/model"
	"github.com/gorilla/mux"
)

//type Server struct {
//	cache *lru.LRUCache
//	router *mux.Router
//}

type Server struct {
	Config *model.ServerConfig
	Router *mux.Router
	Cache model.Cache
}

func NewServer(serverConfig *model.ServerConfig, cache model.Cache) *Server {
	return &Server {
		Config: serverConfig,
		Cache: cache,
	}
}

func (Server *Server) Test1() {

}