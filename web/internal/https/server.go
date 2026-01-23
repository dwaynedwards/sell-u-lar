package https

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"
)

const ShutdownTimeout = 5 * time.Second

type Server struct {
	listener net.Listener
	server   *http.Server
	router   *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: NewRouter(),
	}

	s.server.Handler = http.HandlerFunc(s.router.ServeHTTP)

	return s
}

func (s *Server) Start() error {
	slog.Info("Servering on:", "Host", "localhost", "Port", "3000")
	var err error
	if s.listener, err = net.Listen("tcp", ":3000"); err != nil {
		return err
	}

	go s.server.Serve(s.listener)

	return nil
}

func (s *Server) Stop() error {
	slog.Info("Closing Server on:", "Host", "localhost", "Port", "3000")
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
