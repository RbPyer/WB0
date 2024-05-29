package repository

import "github.com/jmoiron/sqlx"


type OrdersCRUD interface {

}


type Repository struct {
	OrdersCRUD
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}