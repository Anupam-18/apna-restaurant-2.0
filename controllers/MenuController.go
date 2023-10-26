package controllers

import (
	"context"
	"net/http"

	repo "apna-restaurant-2.0/db/sqlc"
	"apna-restaurant-2.0/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MenuController struct {
	db *repo.Queries
}

func NewMenuController(db *repo.Queries) *MenuController {
	return &MenuController{db}
}

func (mc *MenuController) AddMenu(c *gin.Context) {
	var menuReqBody *repo.Menu
	if err := c.ShouldBindJSON(&menuReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddMenuRequest(menuReqBody, mc.db); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	menu := &repo.CreateMenuParams{
		Category:    menuReqBody.Category,
		MenuItemIds: []uuid.UUID{},
	}
	insertedMenu, err := mc.db.CreateMenu(context.Background(), *menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "menu created", "data": insertedMenu})
}

func (mc *MenuController) UpdateMenu(c *gin.Context) {
	var menuReqBody *repo.Menu
	if err := c.ShouldBindJSON(menuReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if resp, ok := utils.ValidateAddMenuRequest(menuReqBody, mc.db); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	menu := &repo.UpdateMenuParams{
		Category:    menuReqBody.Category,
		MenuItemIds: menuReqBody.MenuItemIds,
		ID:          menuReqBody.ID,
	}
	updatedMenu, err := mc.db.UpdateMenu(context.Background(), *menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu updated", "updated_menu": updatedMenu})
}

// TODO - Make this public later
func (mc *MenuController) GetAllMenus(c *gin.Context) {
	allMenus, err := mc.db.GetAllMenus(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menus": allMenus})
}

func (mc *MenuController) GetMenuByID(c *gin.Context) {
	menuId := c.Param("id")
	parsedId, err := uuid.Parse(menuId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menuId"})
		return
	}
	if resp, ok := utils.ValidateMenuItem(parsedId, mc.db); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
	}
	requiredMenu, err := mc.db.GetMenuByID(context.Background(), parsedId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": requiredMenu})
}

func (mc *MenuController) AddMenuItem(c *gin.Context) {

}

func (mc *MenuController) GetAllMenuItems(c *gin.Context) {

}

func (mc *MenuController) GetMenuitemByID(c *gin.Context) {

}

func (mc *MenuController) UpdateMenuitem(c *gin.Context) {

}
