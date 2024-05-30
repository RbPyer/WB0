package server

import (
	"context"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"github.com/RbPyer/WB0/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

type Server struct {
	httpServer *http.Server
	cache map[string]string
	js nats.JetStreamContext
	db repository.Repository

}

func NewServer(port string, handler http.Handler, db *sqlx.DB) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           "127.0.0.1:" + port,
			Handler: handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		db: *repository.NewRepository(db),
		cache: make(map[string]string),
	}
}


func (s *Server) NatsSub(subject string) error{
	conn, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		return err
	}
	js, err := conn.JetStream()
	if err != nil {
		return err
	}

	if _, err := js.Subscribe("Order", func (msg *nats.Msg) {
		log.Println(string(json.RawMessage(msg.Data)))
	}, nats.Durable("wb0")); err != nil {
		return err
	}
	log.Println("Nats JetStream was started...")
	s.js = js
	return nil
}


func (s *Server) Run(subject string) error {
	if err := s.NatsSub(subject); err != nil {
		log.Fatalf("Problems with nats-streaming: %s", err.Error())
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
