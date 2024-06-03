package service

import (
	"github.com/RbPyer/WB0/internal/repository"
	"encoding/json"
)

type OrderService struct {
	repo *repository.Repository
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order_uid string, data json.RawMessage) error {
	return s.repo.Db.CreateOrder(order_uid, data)
}


func (s *OrderService) GetOrders() ([]json.RawMessage, error) {
	return s.repo.Db.GetOrders()
}
