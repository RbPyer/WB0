package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/RbPyer/WB0/internal/service"
	"github.com/RbPyer/WB0/internal/cache"
	"github.com/nats-io/nats.go"
)

type Server struct {
	httpServer *http.Server
	cache *cache.Cache
	js nats.JetStreamContext
	services *service.Service

}

func NewServer(port string, handler http.Handler, services *service.Service, cache *cache.Cache) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           "127.0.0.1:" + port,
			Handler: handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		services: services,
		cache: cache,
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
		log.Printf("A new message in queque:\n\n%s\n\n", string(json.RawMessage(msg.Data)))
		serializedData := make(map[string]interface{})
		err := json.Unmarshal(msg.Data, &serializedData)
		if err != nil {
			log.Fatalf("Some errors while serializing data %s: %s", string(msg.Data), err.Error())
		}
		if err = s.services.CreateOrder(serializedData["order_uid"].(string), msg.Data); err != nil {
			log.Fatalf("Some errors while creating order: %s", err.Error())
		}
		s.cache.Set(serializedData["order_uid"].(string), json.RawMessage(msg.Data))

	}, nats.Durable("wb0")); err != nil {
		return err
	}
	log.Println("Nats JetStream was started...")
	s.js = js
	return nil
}

func (s *Server) CacheLoad() {
	ordersData, err := s.services.GetOrders()
	if err != nil {
		log.Fatalf("Some error while preloading cache from database: %s", err.Error())
	}
	serializedData := make(map[string]interface{})
	for _, record := range ordersData {
		if err = json.Unmarshal(record, &serializedData); err != nil {
			log.Fatalf("Some error while preloading cache with json.Unmarshal(): %s", err.Error())
		}
		s.cache.Set(serializedData["order_uid"].(string), record)
	}
	log.Printf("\n---\nCache was loaded with %d record(s).\n---\n", len(s.cache.Storage))
}


func (s *Server) Run(subject string) error {
	if err := s.NatsSub(subject); err != nil {
		log.Fatalf("Problems with Nats Jetstream: %s", err.Error())
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
