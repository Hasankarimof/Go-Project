package main

import (
	"github.com/gin-gonic/gin"
	"rest-api.com/restapi/db"
	"rest-api.com/restapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
