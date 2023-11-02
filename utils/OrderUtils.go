package utils

import (
	"context"
	"database/sql"
	"fmt"

	repo "apna-restaurant-2.0/db/sqlc"
	"github.com/google/uuid"
)

func ValidateOrderID(id uuid.UUID, db *repo.Queries) (string, bool) {
	orderItemCount, err := db.CheckExistingOrder(context.Background(), id)
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
		for _, orderId := range table.OrderIds {
			if resp, ok := ValidateOrderID(orderId, db); !ok {
				return resp, false
			}
		}
	}
	return "requirement passed", true
}

func ValidateTableId(tableId uuid.UUID, db *repo.Queries) (string, bool) {
	tableCount, err := db.CheckExistingTable(context.Background(), tableId)
	if err != nil {
		return "Internal server error", false
	}
	if tableCount != 1 {
		return "Table does not exist", false
	}
	return "", true
}

func ValidateAddOrderRequest(order *repo.Order, db *repo.Queries, flag string) (string, bool) {
	applyFlag := flag != ""
	if !IsValidUUID(order.ID) && applyFlag {
		return "Missing/Invalid id", false
	} else if resp, ok := ValidateTableId(order.TableID, db); !ok {
		return resp, false
	} else if int(order.Amount.Int32) < 0 {
		return "Amount should be positive", false
	} else if order.CreatedAt.IsZero() {
		return "Created at must be a valid time", false
	} else if order.DeliveredAt.Valid && order.CreatedAt.After(order.DeliveredAt.Time) && applyFlag {
		return "Created time must be earlier than delivered time", false
	}

	for orderItem, quantity := range order.OrderItems.RawMessage {
		orderItemId, err := ConvertIntToUUID(orderItem)
		if err != nil {
			return fmt.Sprintf("Invalid orderitemid %s", orderItemId), false
		}
		if resp, ok := ValidateMenuItem(orderItemId, db); !ok {
			return resp, false
		}
		if quantity < 0 {
			return "One of orderitem's quantity is negative", false
		}
	}
	return "requirement passed", true
}

func CalculateOrderAmount(order *repo.Order, db *repo.Queries) (sql.NullInt32, string) {
	amount := 0
	for orderItem, quantity := range order.OrderItems.RawMessage {
		orderItemId, err := ConvertIntToUUID(orderItem)
		if err != nil {
			return sql.NullInt32{
				Int32: int32(0),
				Valid: true,
			}, "One of orderitems has Invalid orderitem id"
		}
		existingOrderItem, err := db.GetMenuitemsById(context.Background(), orderItemId)
		if err != nil {
			return sql.NullInt32{
				Int32: int32(0),
				Valid: true,
			}, "Internal server error"
		}
		amount += int(existingOrderItem.Price) * int(quantity)
	}
	return sql.NullInt32{
		Int32: int32(amount),
		Valid: true,
	}, "order calculation successful"
}
