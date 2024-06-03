package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RbPyer/WB0/internal/cache"
	"github.com/RbPyer/WB0/internal/service"
	"github.com/RbPyer/WB0/internal/utils"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

type Server struct {
	httpServer *http.Server
	cache *cache.Cache
	nc *nats.Conn
	services *service.Service
	db *sqlx.DB
	crudInterface OrdersCRUD

}

func NewServer(port string, handler http.Handler, services *service.Service, cache *cache.Cache, db *sqlx.DB) *Server {
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
		db: db,
		crudInterface: &service.OrderService{},
	}
}


func (s *Server) NatsSub(subject string) error{
	nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		return err
	}

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	if _, err := js.Subscribe("Order", s.natsHandler, nats.Durable("wb0")); err != nil {
		return err
	}

	log.Println("Nats JetStream was started...")
	s.nc = nc
	return nil
}

func (s *Server) CacheLoad() {
	ordersData, err := GetOrders(s.services.DbService)
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

func (s *Server) natsHandler(msg *nats.Msg) {
	log.Printf("A new message in queque:\n\n%s\n\n", string(json.RawMessage(msg.Data)))
	serializedData := make(map[string]interface{})

	err := json.Unmarshal(msg.Data, &serializedData)
	if err != nil {
		log.Fatalf("Some errors while serializing data %s: %s", string(msg.Data), err.Error())
	}

	if err := utils.ValidateData(serializedData); err != nil {
		log.Fatalf("Some errors with data-validation : %s", err.Error())
	}

	if err := CreateOrder(s.services.DbService, serializedData["order_uid"].(string), msg.Data); err != nil {
		log.Fatalf("Some errors while creating order: %s", err.Error())
	}

	s.cache.Set(serializedData["order_uid"].(string), json.RawMessage(msg.Data))

}

func GracefulShutdown(s *Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func(c chan os.Signal) {
		signal := <-c
		log.Printf("%s was received...", signal.String())
		log.Println("Graceful shutdown was started...")
		if err := s.db.Close(); err != nil {
			log.Printf("Failed to close connection with database: %s\n", err.Error())
		}
		s.nc.Close()
		if err := s.httpServer.Close(); err != nil {
			log.Printf("Failed to close listener in http-server: %s\n", err.Error())
		}
		log.Println("Graceful shutdown was ended...")
	}(c)
}
