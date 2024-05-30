package repository

import (
	"github.com/RbPyer/WB0/internal/models"
	"github.com/jmoiron/sqlx"
)

type OrdersPostgres struct {
	db *sqlx.DB
}


func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{db: db}
}


func (r *OrdersPostgres) CreateOrder(order models.Order) (int, error) {
	return 0, nil
}