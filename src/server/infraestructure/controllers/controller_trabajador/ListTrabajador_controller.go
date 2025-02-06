package controller_trabajador

import (
	"demo/src/server/application/trabajador"
	"demo/src/server/domain/entities"
)

type ListTrabajadorController struct{
	useCase trabajador.ListTrabajador
}

func NewListTrabajadorController(useCase trabajador.ListTrabajador) *ListTrabajadorController {
	return &ListTrabajadorController{useCase: useCase}
}

func (lt_c *ListTrabajadorController) Execute() ([]entities.Trabajador, error) {
	return lt_c.useCase.Execute()
}