package controller_jefeproyecto

import (
	"demo/src/products/application/jefeproyecto"
)

type DeleteJefeProyectoController struct {
	useCase jefeproyecto.DeleteJefeProyecto
}

func NewDeleteJefeProyectoController(useCase jefeproyecto.DeleteJefeProyecto) *DeleteJefeProyectoController {
	return &DeleteJefeProyectoController{useCase: useCase}
}

func (djp_c *DeleteJefeProyectoController) Execute(jefeproyectoID int32) error {
	return djp_c.useCase.Execute(jefeproyectoID)
}