package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"github.com/RbPyer/WB0/pkg/db"
)

type OrdersPostgres struct {
	db *sqlx.DB
}


func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{db: db}
}


func (r *OrdersPostgres) CreateOrder(order_uid string, data json.RawMessage) error {
	query := fmt.Sprintf("INSERT INTO %s (order_uid, order_data) VALUES($1, $2)", db.OrdersTable)
	row := r.db.QueryRow(query, order_uid, data)
	if err := row.Err(); err != nil {
		return err
	}
	log.Printf("A new record with uid <%s> was created in orders!", order_uid)
	return nil
}

func (r *OrdersPostgres) GetOrders() ([]json.RawMessage, error) {
	var response = make([]json.RawMessage, 0)
	query := fmt.Sprintf("SELECT order_data FROM %s", db.OrdersTable)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatalf("Some error while updating cache: %s", err)
	}
	
	for rows.Next() {
		var data json.RawMessage
		if err = rows.Scan(&data); err != nil {
			return nil, err
		}
		response = append(response, data)

	}

	return response, nil

}