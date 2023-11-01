package utils

import (
	"context"

	repo "apna-restaurant-2.0/db/sqlc"
	"github.com/google/uuid"
)

func ValidateOrderItem(item uuid.UUID, db *repo.Queries) (string, bool) {
	orderItemCount, err := db.CheckExistingOrder(context.Background(), item)
	if err != nil {
		return "Internal server error", false
	}
	if orderItemCount != 1 {
		return "One of the Orderitems does not exist", false
	}
	return "", true
}

func ValidateAddTableRequest(table *repo.Table, db *repo.Queries, flag string) (string, bool) {
	if table.TableNumber <= 0 && flag != "update" {
		return "table num required", false
	} else if !IsValidUUID(table.ID) {
		return "Missing/Invalid id", false
	}
	if (len(table.OrderIds)) > 0 {
		for _, orderItem := range table.OrderIds {
			if resp, ok := ValidateOrderItem(orderItem, db); !ok {
				return resp, false
			}
		}
	}
	return "requirement passed", true
}
