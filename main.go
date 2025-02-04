package main

import (
	"demo/src/products/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurar CORS antes de definir las rutas
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config)) // Middleware de CORS

	// Middleware para interceptar las solicitudes OPTIONS
	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Definir rutas
	routes.ProductRoutes(router)
	routes.TrabajadorRoutes(router)
	routes.JefeProyectoRoutes(router)

	// Iniciar el servidor
	router.Run(":8000")
}
