package controllers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"github.com/gin-gonic/gin"
)

type ListTrabajadorController struct {
	useCase useCase_trabajador.ListTrabajador
}

func NewListTrabajadorController(useCase useCase_trabajador.ListTrabajador) *ListTrabajadorController {
	return &ListTrabajadorController{useCase: useCase}
}

func (c *ListTrabajadorController) Execute(ctx *gin.Context) {
	trabajadores, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	converted := make([]map[string]interface{}, 0)
	for _, t := range trabajadores {
		convertedTrabajador := map[string]interface{}{
			"idtrabajador":     t.Idtrabajador,
			"nombretrabajador": t.Nombretrabajador,
			"posicion":         t.Posicion,
			"telefono":         t.Telefono,
			"correo":           t.Correo,
			"salario":          t.Salario,
			"aniosexperiencia": t.Aniosexperiencia,
		}
		converted = append(converted, convertedTrabajador)
	}

	ctx.JSON(200, converted)
}
