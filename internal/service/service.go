package service

import (
	"github.com/RbPyer/WB0/internal/models"
	"github.com/RbPyer/WB0/internal/repository"
)

type OrdersCRUD interface {
	CreateOrder(order models.Order) (int, error)
}


type Service struct {
	OrdersCRUD
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		OrdersCRUD: NewOrderService(repo.OrdersCRUD),
	}
}