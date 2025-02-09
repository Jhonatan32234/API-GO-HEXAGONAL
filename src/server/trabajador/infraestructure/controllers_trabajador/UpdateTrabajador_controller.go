package controllers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"demo/src/server/trabajador/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateTrabajadorController struct {
	useCase useCase_trabajador.UpdateTrabajador
}

func NewUpdateTrabajadorController(useCase useCase_trabajador.UpdateTrabajador) *UpdateTrabajadorController {
	return &UpdateTrabajadorController{useCase: useCase}
}

func (ut_c *UpdateTrabajadorController) Execute(c *gin.Context) {
	var trabajador entities.Trabajador

	if err := c.ShouldBindJSON(&trabajador); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	trabajador.Idtrabajador = int32(id)

	if err := ut_c.useCase.Execute(trabajador); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Trabajador actualizado correctamente"})
}
