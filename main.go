package main

import (
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterUserRoutes(r)
	r.Run(":8080") // Servidor corriendo en el puerto 8080
}
