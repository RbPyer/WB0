package repository

import (
	"github.com/RbPyer/WB0/internal/models"
	"github.com/jmoiron/sqlx"
)


type OrdersCRUD interface {
	CreateOrder(order models.Order) (int, error)
}


type Repository struct {
	OrdersCRUD
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrdersCRUD: NewOrdersPostgres(db),
	}
}