package app

import "apna-restaurant-2.0/routes"

func mapUrls(authRoute *routes.AuthRoutes, menuRoute *routes.MenuRoutes) {
	authRoute.AuthRoute(&router.RouterGroup)
	menuRoute.MenuRoute(&router.RouterGroup)
}
