package service

import (
	"github.com/RbPyer/WB0/internal/models"
	"github.com/RbPyer/WB0/internal/repository"
)

type OrderService struct {
	repo repository.OrdersCRUD
}

func NewOrderService(repo repository.OrdersCRUD) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order models.Order) (int, error) {
	return s.repo.CreateOrder(order)
}

