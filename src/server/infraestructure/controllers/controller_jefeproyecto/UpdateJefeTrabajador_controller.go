package controller_jefeproyecto

import (
	"demo/src/server/application/jefeproyecto"
	"demo/src/server/domain/entities"
)
type UpdateJefeProyectoController struct {
	useCase jefeproyecto.UpdateJefeProyecto
}

func NewUpdateJefeProyectoController(useCase jefeproyecto.UpdateJefeProyecto) *UpdateJefeProyectoController {
	return &UpdateJefeProyectoController{useCase: useCase}
}

func (ujp_c *UpdateJefeProyectoController) Execute(jefeproyecto entities.JefeProyecto) error {
	return ujp_c.useCase.Execute(jefeproyecto)
}
