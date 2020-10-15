package server

import (
	"fmt"
	"github.com/bhupeshpandey/cars/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Config *model.ServerConfig
	Router *mux.Router
	Cache  model.Cache
}

func New(serverConfig *model.ServerConfig, cache model.Cache, db model.DB) *Server {
	router := mux.NewRouter()

	server := &Server{
		Config: serverConfig,
		Cache:  cache,
		Router: router,
	}

	router.HandleFunc("/api/cars", server.GetCars).Methods("GET")
	router.HandleFunc("/api/car/{id}", server.GetCarWithId).Methods("GET")
	router.HandleFunc("/api/accessories", server.GetAccessories).Methods("GET")
	router.HandleFunc("/api/books/{id}", server.GetAccessoryWithId).Methods("GET")
	router.HandleFunc("/api/recent/items}", server.GetRecentItems).Methods("GET")
	router.HandleFunc("/api/recent/items}", server.SetSelectedItem).Methods("PUT")
	// set our port address

	return server
}

func (Server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", Server.Config.Host, Server.Config.Port), Server.Router))
}

//func (ServerConfig *ServerConfig) GetProducts()

func (Server *Server) GetCars(w http.ResponseWriter, r *http.Request) {

}
func (Server *Server) GetCarWithId(w http.ResponseWriter, r *http.Request) {

}

func (Server *Server) GetAccessories(w http.ResponseWriter, r *http.Request) {

}

func (Server *Server) GetAccessoryWithId(w http.ResponseWriter, r *http.Request) {

}

func (Server *Server) GetRecentItems(w http.ResponseWriter, r *http.Request) {

}

func (Server *Server) SetSelectedItem(w http.ResponseWriter, r *http.Request) {

}

func (Server *Server) SetSelectedCar(w http.ResponseWriter, r *http.Request) {

}
