package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type OrdersPostgres struct {
	db *sqlx.DB
}


func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{db: db}
}


func (r *OrdersPostgres) CreateOrder(order_uid string, data json.RawMessage) error {
	query := fmt.Sprintf("INSERT INTO %s (order_uid, order_data) VALUES($1, $2)", ordersTable)
	row := r.db.QueryRow(query, order_uid, data)
	if err := row.Err(); err != nil {
		return err
	}
	log.Printf("A new record with uid <%s> was created in orders!", order_uid)
	return nil
}