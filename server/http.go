package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mfojtik/bugtraq/services"
)

type HttpServer struct {
	Port string
	Host string

	service services.Service
}

func NewServer(bindHost, port string) *HttpServer {
	return &HttpServer{Host: bindHost, Port: port}
}

func (s *HttpServer) ListHandler(w http.ResponseWriter, r *http.Request) {

	var (
		jsonResponse string
		err          error
	)

	if jsonResponse, err = s.service.List(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, jsonResponse)
}

func (s *HttpServer) SetService(service services.Service) {
	s.service = service
}

func (b *HttpServer) Start() {
	log.Printf("Started HTTP server http://%s:%s\n", b.Host, b.Port)
	http.HandleFunc("/feed", b.ListHandler)
	http.ListenAndServe(b.Host+":"+b.Port, nil)
}
