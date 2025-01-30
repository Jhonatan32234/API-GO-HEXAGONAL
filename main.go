package main

import (
	"demo/src/products/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.ProductRoutes(router)
	routes.TrabajadorRoutes(router)
	routes.JefeProyectoRoutes(router)
	router.Run(":8000")
}