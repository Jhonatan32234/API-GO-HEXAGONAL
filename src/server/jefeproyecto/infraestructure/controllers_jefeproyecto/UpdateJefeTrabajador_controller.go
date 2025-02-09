package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateJefeProyectoController struct {
	useCase useCase_jefeproyecto.UpdateJefeProyecto
}

func NewUpdateJefeProyectoController(useCase useCase_jefeproyecto.UpdateJefeProyecto) *UpdateJefeProyectoController {
	return &UpdateJefeProyectoController{useCase: useCase}
}

func (c *UpdateJefeProyectoController) Execute(ctx *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	// Validar el JSON recibido
	if err := ctx.ShouldBindJSON(&jefeproyectos); err != nil {
		ctx.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	// Obtener y validar el ID de la URL
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// Asignar el ID al objeto JefeProyecto
	jefeproyectos.Idjefeproyecto = int32(id)

	// Ejecutar la actualización
	if err := c.useCase.Execute(jefeproyectos); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Responder con éxito
	ctx.JSON(200, gin.H{"message": "JefeProyecto actualizado correctamente"})
}
