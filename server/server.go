package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/medhir/yaml-api/storage"
)

type Server struct {
	ctx     context.Context
	router  *http.ServeMux
	server  *http.Server
	storage *storage.Storage
}

// NewServer initializes a server object
func NewServer(port string) *Server {
	server := &Server{
		ctx:    context.Background(),
		router: http.DefaultServeMux,
		server: &http.Server{
			Addr: port,
		},
		storage: storage.NewStorage(),
	}
	server.setRoutes()
	return server
}

// Start opens a connection to accept http traffic on the desired port
func (s *Server) Start() {
	fmt.Printf("Listening on port %s\n", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println("Server stopped unexpectedly:", err)
		s.shutdown()
	}
}

// shutdown gracefully shuts down a server instance
func (s *Server) shutdown() {
	if s.server != nil {
		ctx, cancel := context.WithTimeout(s.ctx, 10*time.Second)
		defer cancel()
		err := s.server.Shutdown(ctx)
		if err != nil {
			fmt.Println("Failed to shut down server gracefully.", err)
		}
	}
}
