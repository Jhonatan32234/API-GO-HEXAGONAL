package handlers_jefeproyecto



import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/controllers_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/models_jefeproyecto"
	"github.com/gin-gonic/gin"
)

func ListJefeProyectosHandler(c *gin.Context) {
	ps := models_jefeproyecto.NewMySQLJP()
	useCase := useCase_jefeproyecto.NewListJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewListJefeProyectoController(*useCase)

	controller.Execute(c) // Delegación de la ejecución al controlador
}