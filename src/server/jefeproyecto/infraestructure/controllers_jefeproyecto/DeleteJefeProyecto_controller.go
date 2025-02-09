package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteJefeProyectoController struct {
	useCase useCase_jefeproyecto.DeleteJefeProyecto
}

func NewDeleteJefeProyectoController(useCase useCase_jefeproyecto.DeleteJefeProyecto) *DeleteJefeProyectoController {
	return &DeleteJefeProyectoController{useCase: useCase}
}

func (c *DeleteJefeProyectoController) Execute(ctx *gin.Context) {
	// Obtener y validar el ID de la URL
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar la eliminación
	if err := c.useCase.Execute(int32(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Responder con éxito
	ctx.JSON(200, gin.H{"message": "JefeProyecto eliminado correctamente"})
}
