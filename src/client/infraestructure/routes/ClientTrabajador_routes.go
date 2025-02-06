package routes

import (
    "github.com/gin-gonic/gin"
    "demo/src/client/application/trabajador"
    "demo/src/client/infraestructure/controllers/trabajador_controller"
)
var apiURLTrabajador = "http://localhost:8000" 


func TrabjadorRoutes(router *gin.Engine) {
    router.GET("/trabajador/salario/:id", checkSalaryChangeHandler)
    router.GET("/trabajador/posicion/:id", checkPositionChangeHandler)
}

func checkSalaryChangeHandler(c *gin.Context) {
   
    useCase := trabajador.NewCheckSalaryChange(apiURLTrabajador)
    controller := trabajador_controller.NewCheckSalaryChangeController(*useCase)
    controller.Execute(c)
}

func checkPositionChangeHandler(c *gin.Context) {
   
    useCase := trabajador.NewCheckpositionChange(apiURLTrabajador)
    controller := trabajador_controller.NewCheckPositionChangeController(*useCase)
    controller.Execute(c)
}