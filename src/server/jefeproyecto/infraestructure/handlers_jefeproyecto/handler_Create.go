package handlers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/controllers_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/models_jefeproyecto"

	"github.com/gin-gonic/gin"
)

func CreateJefeProyectoHandler(c *gin.Context) {
	ps := models_jefeproyecto.NewMySQLJP()
	useCase := useCase_jefeproyecto.NewCreateJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewCreateJefeProyectoController(*useCase)
	controller.Execute(c)
}