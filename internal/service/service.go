package service

import (
	"github.com/RbPyer/WB0/internal/repository"
)




type Service struct {
	DbService *OrderService
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		DbService: NewOrderService(repo),
	}
}