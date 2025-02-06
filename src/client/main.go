package client


import (
    "github.com/gin-gonic/gin"
    "demo/src/client/infraestructure/routes"
)

func Run() {
    router := gin.Default()
    routes.TrabjadorRoutes(router)
    routes.JefeProyectoRoutes(router)
    router.Run(":8081") // P
	
}