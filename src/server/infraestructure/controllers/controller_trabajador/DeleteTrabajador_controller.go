package controller_trabajador

import (
	"demo/src/server/application/trabajador"
)

type DeleteTrabajadorController struct {
	useCase trabajador.DeleteTrabajador
}

func NewDeleteTrabajadorController(useCase trabajador.DeleteTrabajador) *DeleteTrabajadorController {
	return &DeleteTrabajadorController{useCase: useCase}
}

func (dt_c *DeleteTrabajadorController) Execute(trabajadorID int32) error {
	return dt_c.useCase.Execute(trabajadorID)
}