package repository

import (
	"github.com/jmoiron/sqlx"
)


type Repository struct {
	Db *OrdersPostgres
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: NewOrdersPostgres(db),
	}
}