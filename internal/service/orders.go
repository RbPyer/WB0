package service

import (
	"github.com/RbPyer/WB0/internal/repository"
	"encoding/json"
)

type OrderService struct {
	repo repository.OrdersCRUD
}

func NewOrderService(repo repository.OrdersCRUD) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order_uid string, data json.RawMessage) error {
	return s.repo.CreateOrder(order_uid, data)
}

