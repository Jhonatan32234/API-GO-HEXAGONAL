package controller_trabajador

import (
	"demo/src/products/application/trabajador"
	"demo/src/products/domain/entities"
)
type UpdateTrabajadorController struct {
	useCase trabajador.UpdateTrabajador
}

func NewUpdateTrabajadorController(useCase trabajador.UpdateTrabajador) *UpdateTrabajadorController {
	return &UpdateTrabajadorController{useCase: useCase}
}

func (ut_c *UpdateTrabajadorController) Execute(trabajador entities.Trabajador) error {
	return ut_c.useCase.Execute(trabajador)
}
