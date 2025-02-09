package handlers_trabajador


import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"demo/src/server/trabajador/infraestructure/controllers_trabajador"
	"demo/src/server/trabajador/infraestructure/models_trabajador"
	"github.com/gin-gonic/gin"
)

func DeleteTrabajadorHandler(c *gin.Context) {
	ps := models_trabajador.NewMySQLT()
	useCase := useCase_trabajador.NewDeleteTrabajador(ps)
	controller := controllers_trabajador.NewDeleteTrabajadorController(*useCase)

	controller.Execute(c) 
}
