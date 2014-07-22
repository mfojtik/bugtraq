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

func NewServer(bindHost, port string, s services.Service) *HttpServer {
	return &HttpServer{Host: bindHost, Port: port, service: s}
}

func (s *HttpServer) JsonListHandler(w http.ResponseWriter, r *http.Request) {
	if jsonResponse, err := s.service.GetListJSON(); err == nil {
		w.Header().Set("X-Last-Update", s.service.LastUpdate().String())
		fmt.Fprintf(w, jsonResponse)
	} else {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (b *HttpServer) Start() {
	log.Printf("Started HTTP server http://%s:%s/%s\n", b.Host, b.Port, b.service.Id())
	http.HandleFunc(fmt.Sprintf("/%s", b.service.Id()), b.JsonListHandler)
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	http.ListenAndServe(b.Host+":"+b.Port, nil)
}
