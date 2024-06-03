package utils

import (
	"errors"
	"slices"
	"log"
)


func ValidateData(data map[string]any) error {
	var required_params []string = []string{
		"order_uid",
		"track_number",
		"entry",
		"delivery",
		"payment",
		"items",
		"locale",
		"internal_signature",
		"customer_id",
		"delivery_service",
		"shardkey",
		"sm_id",
		"date_created",
		"oof_shard",
	}
	if !checkKeys(data, required_params) {
		return errors.New("not enough required params")
	}

	return nil
}


func checkKeys(data map[string]any, dataset []string) bool {
	counter := 0
	for k := range data {
		if slices.Contains(dataset, k) && data[k] != nil {
			counter++
		}
	}

	log.Println(counter)
	return counter == len(dataset)

}