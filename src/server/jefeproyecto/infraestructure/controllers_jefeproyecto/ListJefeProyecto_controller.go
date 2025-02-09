


/*package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/domain/entities"
)

type ListJefeProyectoController struct{
	useCase useCase_jefeproyecto.ListJefeProyecto
}

func NewListJefeProyectoController(useCase useCase_jefeproyecto.ListJefeProyecto) *ListJefeProyectoController {
	return &ListJefeProyectoController{useCase: useCase}
}

func (ljp_c *ListJefeProyectoController) Execute() ([]entities.JefeProyecto, error) {
	return ljp_c.useCase.Execute()
}*/


package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"github.com/gin-gonic/gin"
)

type ListJefeProyectoController struct {
	useCase useCase_jefeproyecto.ListJefeProyecto
}

func NewListJefeProyectoController(useCase useCase_jefeproyecto.ListJefeProyecto) *ListJefeProyectoController {
	return &ListJefeProyectoController{useCase: useCase}
}

func (c *ListJefeProyectoController) Execute(ctx *gin.Context) {
	// Obtener la lista de Jefes de Proyecto desde el caso de uso
	jefeproyectos, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Convertir la lista de Jefes de Proyecto a un formato adecuado
	converted := make([]map[string]interface{}, len(jefeproyectos))
	for i, jp := range jefeproyectos {
		converted[i] = map[string]interface{}{
			"idjefeproyecto":   jp.Idjefeproyecto,
			"nombrejefe":       jp.Nombrejefe,
			"telefono":         jp.Telefono,
			"correo":           jp.Correo,
			"salario":          jp.Salario,
			"aniosexperiencia": jp.Aniosexperiencia,
		}
	}

	// Responder con la lista formateada
	ctx.JSON(200, converted)
}
