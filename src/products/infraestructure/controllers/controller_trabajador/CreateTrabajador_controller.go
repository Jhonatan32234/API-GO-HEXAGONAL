package controller_trabajador

import (
	"demo/src/products/application/trabajador"
	"demo/src/products/domain/entities"
)


type CreateTrabajadorController struct{
	useCase trabajador.CreateTrabajador
}

func NewCreateTrabajadorController(useCase trabajador.CreateTrabajador) *CreateTrabajadorController {
	return &CreateTrabajadorController{useCase: useCase}
}

func (ct_c *CreateTrabajadorController) Execute(trabajador entities.Trabajador) error{
	return ct_c.useCase.Execute(trabajador)
}