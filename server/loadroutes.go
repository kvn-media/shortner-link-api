package server

import "github.com/kvn-media/link-shortner-api/server/routes"

func loadRoutes() {
	ServerObject.GET("link", routes.GetRedirect)
	ServerObject.GET("new", routes.Shorten)
}
