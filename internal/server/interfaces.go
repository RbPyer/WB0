package server

import "encoding/json"

type OrdersCRUD interface {
	CreateOrder(order_uid string, data json.RawMessage) error
	GetOrders() ([]json.RawMessage, error)
}


func CreateOrder(oc OrdersCRUD, order_uid string, data json.RawMessage) error {
	return oc.CreateOrder(order_uid, data)
}


func GetOrders(oc OrdersCRUD) ([]json.RawMessage, error) {
	return oc.GetOrders()
}