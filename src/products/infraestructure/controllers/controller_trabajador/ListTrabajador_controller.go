package controller_trabajador

import (
	"demo/src/products/application/trabajador"
	"demo/src/products/domain/entities"
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