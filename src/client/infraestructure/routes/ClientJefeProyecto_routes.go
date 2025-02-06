package routes

import (
	"demo/src/client/application/jefeproyecto"
	"demo/src/client/infraestructure/controllers/jefeproyecto_controller"

	"github.com/gin-gonic/gin"
)

var apiURLJefe = "http://localhost:8000"

func JefeProyectoRoutes(router *gin.Engine) {
	router.GET("/jefeproyecto/experiencia/:id", checkExperienceChangeHandler)
}

func checkExperienceChangeHandler(c *gin.Context) {
	useCase := jefeproyecto.NewCheckExperienceChange(apiURLJefe)
	controller := jefeproyecto_controller.NewCheckExperienceChangeController(*useCase)
	controller.Execute(c)
}