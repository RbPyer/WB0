package service

import "github.com/RbPyer/WB0/pkg/repository"

type OrdersCRUD interface {

}


type Service struct {
	OrdersCRUD
}


func NewService(repo *repository.Repository) *Service {
	return &Service{}
}