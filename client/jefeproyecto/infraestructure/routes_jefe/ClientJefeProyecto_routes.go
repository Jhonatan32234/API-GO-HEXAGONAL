package routes_jefe

import (
	"demo/client/jefeproyecto/application/useCase_jefe"
	"demo/client/jefeproyecto/infraestructure/controllers_jefe"

	"github.com/gin-gonic/gin"
)

var apiURLJefe = "http://localhost:8000"

func JefeProyectoRoutes(router *gin.Engine) {
	useCase := useCase_jefe.NewCheckExperienceChange(apiURLJefe)
	controller := controllers_jefe.NewCheckExperienceChangeController(*useCase)

	router.GET("/jefeproyecto/experiencia/:id", controller.Execute)
}
