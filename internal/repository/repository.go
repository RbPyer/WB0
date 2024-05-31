package repository

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
)


type OrdersCRUD interface {
	CreateOrder(order_uid string, data json.RawMessage) error
}


type Repository struct {
	OrdersCRUD
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrdersCRUD: NewOrdersPostgres(db),
	}
}