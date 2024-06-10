package main

import (
	"sample/router"
	"sample/server"
)

func main() {
	routes := router.SetupRoutes()
	server.StartServer(routes)
}
