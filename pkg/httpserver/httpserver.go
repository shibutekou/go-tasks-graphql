package httpserver

import (
	"context"
	"github.com/vgekko/go-tasks-graphql/config"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, cfg config.HTTP) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    cfg.Addr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.ShutdownTimeout,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
