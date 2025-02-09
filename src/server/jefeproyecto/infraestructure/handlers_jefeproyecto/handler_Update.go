package handlers_jefeproyecto


import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/controllers_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/models_jefeproyecto"
	"github.com/gin-gonic/gin"
)

func UpdateJefeProyectosHandler(c *gin.Context) {
	ps := models_jefeproyecto.NewMySQLJP()
	useCase := useCase_jefeproyecto.NewUpdateJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewUpdateJefeProyectoController(*useCase)

	controller.Execute(c) 
}