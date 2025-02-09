package controllers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"demo/src/server/trabajador/domain/entities"
	"github.com/gin-gonic/gin"
	"fmt"
)

type CreateTrabajadorController struct {
	useCase useCase_trabajador.CreateTrabajador
}

func NewCreateTrabajadorController(useCase useCase_trabajador.CreateTrabajador) *CreateTrabajadorController {
	return &CreateTrabajadorController{useCase: useCase}
}

func (c *CreateTrabajadorController) Execute(ctx *gin.Context) {
	var trabajador entities.Trabajador

	if err := ctx.ShouldBindJSON(&trabajador); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		ctx.JSON(400, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	if err := c.useCase.Execute(trabajador); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Trabajador creado exitosamente",
		"trabajador": trabajador,
	})
}
