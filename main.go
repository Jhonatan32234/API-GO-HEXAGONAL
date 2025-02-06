package main

import (
	//"demo/src/client"
	"demo/src/client"
	"demo/src/server/infraestructure/routes"
	//"os"
	//"os/exec"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config)) 

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

	routes.ProductRoutes(router)
	routes.TrabajadorRoutes(router)
	routes.JefeProyectoRoutes(router)

	go func() {
        time.Sleep(2 * time.Second) // Esperar 2 segundos para asegurarse de que el servidor esté levantado
        client.Run()
    }()

	
	router.Run(":8000")
}
