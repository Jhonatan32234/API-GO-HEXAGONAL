package routes_trabajador

import (
	"demo/client/trabajador/application/useCase_trabajador"
	"demo/client/trabajador/infraestructure/controllers_trabajador"
	"github.com/gin-gonic/gin"
)
var apiURLTrabajador = "http://localhost:8000" 


func TrabjadorRoutes(router *gin.Engine) {
    router.GET("/trabajador/salario/:id", checkSalaryChangeHandler)
    router.GET("/trabajador/posicion/:id", checkPositionChangeHandler)
}

func checkSalaryChangeHandler(c *gin.Context) {
   
    useCase := useCase_trabajador.NewCheckSalaryChange(apiURLTrabajador)
    controller := controllers_trabajador.NewCheckSalaryChangeController(*useCase)
    controller.Execute(c)
}

func checkPositionChangeHandler(c *gin.Context) {
   
    useCase := useCase_trabajador.NewCheckpositionChange(apiURLTrabajador)
    controller := controllers_trabajador.NewCheckPositionChangeController(*useCase)
    controller.Execute(c)
}