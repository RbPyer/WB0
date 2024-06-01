package service

import (
	"github.com/RbPyer/WB0/internal/repository"
	"encoding/json"
)

type OrdersCRUD interface {
	CreateOrder(order_uid string, data json.RawMessage) error
	GetOrders() ([]json.RawMessage, error)
}


type Service struct {
	OrdersCRUD
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		OrdersCRUD: NewOrderService(repo.OrdersCRUD),
	}
}