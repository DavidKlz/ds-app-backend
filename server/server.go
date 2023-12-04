package server

import (
	"log"
	"net/http"

	"github.com/davidklz/ds-app-backend/database"
	"github.com/gorilla/mux"
)

type Server struct {
	port  string
	store database.Store
}

func InitializeServer(port string, store database.Store) *Server {
	return &Server{
		port:  port,
		store: store,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/formular", createHttpHandlerFunc(s.handleFormularRequests))
	router.HandleFunc("/variable", createHttpHandlerFunc(s.handleVariableRequests))
	router.HandleFunc("/datentyp", createHttpHandlerFunc(s.handleDatentypRequests))
	router.HandleFunc("/controltyp", createHttpHandlerFunc(s.handleControltypRequests))

	router.HandleFunc("/formular/{id}", createHttpHandlerFunc(s.handleFormularRequestsById))
	router.HandleFunc("/variable/{id}", createHttpHandlerFunc(s.handleVariableRequestsById))
	router.HandleFunc("/datentyp/{id}", createHttpHandlerFunc(s.handleDatentypRequestsById))
	router.HandleFunc("/controltyp/{id}", createHttpHandlerFunc(s.handleControltypRequestsById))

	log.Println("Server is running on port: ", s.port)

	http.ListenAndServe(s.port, router)
}
