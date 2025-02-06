package controller_trabajador

import (
	"demo/src/server/application/trabajador"
	"demo/src/server/domain/entities"
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