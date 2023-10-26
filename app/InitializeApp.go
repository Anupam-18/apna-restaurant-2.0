package app

import (
	"apna-restaurant-2.0/controllers"
	"apna-restaurant-2.0/db/config"
	repo "apna-restaurant-2.0/db/sqlc"
	"apna-restaurant-2.0/routes"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	db     *repo.Queries
	// AuthController *controllers.AuthController
	// AuthRoutes     *routes.AuthRoutes
)

func init() {
	// db := config.ConnectToDB()
	// router = gin.Default()
	// queries := repo.New(db)
	// AuthController = controllers.NewAuthController(queries)
	// AuthRoutes = routes.NewAuthRoutes(AuthController, router)
}

func StartApplication() {
	db := config.ConnectToDB()
	router = gin.Default()
	queries := repo.New(db)
	AuthController := controllers.NewAuthController(queries)
	AuthRoutes := routes.NewAuthRoutes(AuthController, router)
	MenuController := controllers.NewMenuController(queries)
	MenuRoutes := routes.NewMenuRoutes(MenuController, router)
	mapUrls(AuthRoutes, MenuRoutes)
	router.Run(":8080")
}
