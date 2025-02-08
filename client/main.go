package client

import (
	"demo/client/jefeproyecto/infraestructure/routes_jefe"
	"demo/client/trabajador/infraestructure/routes_trabajador"

	"github.com/gin-gonic/gin"
)

func Run() {
    router := gin.Default()
    routes_trabajador.TrabjadorRoutes(router)
    routes_jefe.JefeProyectoRoutes(router)
    router.Run(":8081") // :P  :) :D :O :/  :|  :(  >:(  
	
}