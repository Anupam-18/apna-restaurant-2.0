package controllers

import (
	"context"
	"net/http"

	repo "apna-restaurant-2.0/db/sqlc"
	"apna-restaurant-2.0/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	db *repo.Queries
}

func NewOrderController(db *repo.Queries) *OrderController {
	return &OrderController{db: db}
}

func (oc *OrderController) AddTable(c *gin.Context) {
	var tableReqBody *repo.Table
	if err := c.ShouldBindJSON(&tableReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddTableRequest(tableReqBody, oc.db, ""); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	order := &repo.CreateTableParams{
		TableNumber: tableReqBody.TableNumber,
		OrderIds:    tableReqBody.OrderIds,
	}
	insertedTable, err := oc.db.CreateTable(context.Background(), *order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting table in db"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "table created", "data": insertedTable})
}

// TODO: while testing, check for invalid table_id,w/t giving table_id
func (oc *OrderController) UpdateTable(c *gin.Context) {
	var tableReqBody *repo.Table
	if err := c.ShouldBindJSON(&tableReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddTableRequest(tableReqBody, oc.db, "update"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	existingTable, err := oc.db.GetTableByID(context.Background(), tableReqBody.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if existingTable.TableNumber == 0 {
		tableReqBody.TableNumber = existingTable.TableNumber
	}
	if len(existingTable.OrderIds) == 0 {
		tableReqBody.OrderIds = existingTable.OrderIds
	}
	table := &repo.UpdateTableParams{
		TableNumber: existingTable.TableNumber,
		OrderIds:    existingTable.OrderIds,
		ID:          existingTable.ID,
	}
	updatedTable, err := oc.db.UpdateTable(context.Background(), *table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table updated", "updated_table": updatedTable})
}

func (oc *OrderController) AddOrder(c *gin.Context) {
	var orderReqBody *repo.Order
	if err := c.ShouldBindJSON(&orderReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddOrderRequest(orderReqBody, oc.db, ""); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	amount, msg := utils.CalculateOrderAmount(orderReqBody, oc.db)
	if msg != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	order := &repo.CreateOrderParams{
		TableID:    orderReqBody.TableID,
		Amount:     amount,
		OrderItems: orderReqBody.OrderItems,
	}
	insertedOrder, err := oc.db.CreateOrder(context.Background(), *order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created", "data": insertedOrder})
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
	var orderReqBody *repo.Order
	if err := c.ShouldBindJSON(&orderReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddOrderRequest(orderReqBody, oc.db, "update"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}

	existingOrder, err := oc.db.GetOrderByID(context.Background(), orderReqBody.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if existingOrder.TableID != orderReqBody.TableID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table id cannot be changed"})
		return
	}
	if !orderReqBody.OrderItems.Valid {
		orderReqBody.OrderItems = existingOrder.OrderItems
	}
	amount, msg := utils.CalculateOrderAmount(orderReqBody, oc.db)
	if msg != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if orderReqBody.CreatedAt.IsZero() {
		orderReqBody.CreatedAt = existingOrder.CreatedAt
	}
	table := &repo.UpdateOrderParams{
		TableID:     orderReqBody.TableID,
		Amount:      amount,
		OrderItems:  orderReqBody.OrderItems,
		CreatedAt:   orderReqBody.CreatedAt,
		DeliveredAt: orderReqBody.DeliveredAt,
		ID:          orderReqBody.ID,
	}
	updatedOrder, err := oc.db.UpdateOrder(context.Background(), *table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated", "updated_order": updatedOrder})
}

func (oc *OrderController) GetOrderDetails(c *gin.Context) {
	orderId := c.Param("id")
	parsedId, err := uuid.Parse(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid orderId"})
		return
	}
	if resp, ok := utils.ValidateOrderID(parsedId, oc.db); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	requiredOrder, err := oc.db.GetOrderByID(context.Background(), parsedId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": requiredOrder})
}

func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderId := c.Param("id")
	parsedOrderId, err := uuid.Parse(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tableId"})
		return
	}
	err = oc.db.DeleteOrderByID(context.Background(), parsedOrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting order from db"})
		return
	}
	err = oc.db.RemoveOrderIDFromTables(context.Background(), parsedOrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting order from db"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})

}

func (oc *OrderController) GetOrderDetailsForTable(c *gin.Context) {
	tableId := c.Param("table_id")
	parsedTableId, err := uuid.Parse(tableId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tableId"})
		return
	}
	if resp, ok := utils.ValidateTableId(parsedTableId, oc.db); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	requiredOrders := make([]repo.Order, 0)

	existingTable, err := oc.db.GetTableByID(context.Background(), parsedTableId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	for _, orderId := range existingTable.OrderIds {
		requiredOrder, err := oc.db.GetOrderByID(context.Background(), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		requiredOrders = append(requiredOrders, requiredOrder)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orders fetched", "data": requiredOrders})
}
