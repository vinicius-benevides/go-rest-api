package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/routes"
)

func main() {
	server := gin.Default()
	if err := db.Init(); err != nil {
		panic(err)
	}

	routes.Register(server)

	server.Run(":8080")
}
