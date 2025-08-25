package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-crud-app/db"
	"golang-crud-app/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
