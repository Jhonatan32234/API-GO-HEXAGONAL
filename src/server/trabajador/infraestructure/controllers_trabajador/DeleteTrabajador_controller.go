package controllers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteTrabajadorController struct {
	useCase useCase_trabajador.DeleteTrabajador
}

func NewDeleteTrabajadorController(useCase useCase_trabajador.DeleteTrabajador) *DeleteTrabajadorController {
	return &DeleteTrabajadorController{useCase: useCase}
}

func (c *DeleteTrabajadorController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(int32(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Trabajador eliminado correctamente"})
}
