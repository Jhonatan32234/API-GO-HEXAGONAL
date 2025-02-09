

/*package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/domain/entities"
)


type CreateJefeProyectoController struct{
	useCase useCase_jefeproyecto.CreateJefeProyecto
}

func NewCreateJefeProyectoController(useCase useCase_jefeproyecto.CreateJefeProyecto) *CreateJefeProyectoController {
	return &CreateJefeProyectoController{useCase: useCase}
}

func (cjp_c *CreateJefeProyectoController) Execute(jefeproyecto entities.JefeProyecto) error{
	return cjp_c.useCase.Execute(jefeproyecto)
}*/
package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/domain/entities"
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"github.com/gin-gonic/gin"
	"fmt"
)

type CreateJefeProyectoController struct {
	useCase useCase_jefeproyecto.CreateJefeProyecto
}

func NewCreateJefeProyectoController(useCase useCase_jefeproyecto.CreateJefeProyecto) *CreateJefeProyectoController {
	return &CreateJefeProyectoController{useCase: useCase}
}

func (c *CreateJefeProyectoController) Execute(ctx *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	// Validación de datos
	if err := ctx.ShouldBindJSON(&jefeproyectos); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		ctx.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	// Ejecución del caso de uso
	if err := c.useCase.Execute(jefeproyectos); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta al cliente
	ctx.JSON(201, gin.H{
		"message": "JefeProyecto creado exitosamente",
		"jefeproyecto": jefeproyectos, 
	})
}
