package controllers

import (
	"context"
	"net/http"

	repo "apna-restaurant-2.0/db/sqlc"
	"apna-restaurant-2.0/utils"
	"github.com/gin-gonic/gin"
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

}

func (oc *OrderController) UpdateOrder(c *gin.Context) {

}

func (oc *OrderController) GetOrderDetails(c *gin.Context) {

}

func (oc *OrderController) CancelOrder(c *gin.Context) {

}

func (oc *OrderController) GetOrderDetailsForTable(c *gin.Context) {

}
