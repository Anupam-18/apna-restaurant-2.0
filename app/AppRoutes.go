package app

import "apna-restaurant-2.0/routes"

func mapUrls(authRoute *routes.AuthRoutes) {
	authRoute.AuthRoute(&router.RouterGroup)
}
