package handlers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"demo/src/server/trabajador/infraestructure/controllers_trabajador"
	"demo/src/server/trabajador/infraestructure/models_trabajador"
	"github.com/gin-gonic/gin"
)

func ListTrabajadoresHandler(c *gin.Context) {
	ps := models_trabajador.NewMySQLT()
	useCase := useCase_trabajador.NewListTrabajador(ps)
	controller := controllers_trabajador.NewListTrabajadorController(*useCase)

	controller.Execute(c) // Delegación de la ejecución al controlador
}