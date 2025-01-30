package controller_jefeproyecto

import (
	"demo/src/products/application/jefeproyecto"
	"demo/src/products/domain/entities"
)


type CreateJefeProyectoController struct{
	useCase jefeproyecto.CreateJefeProyecto
}

func NewCreateJefeProyectoController(useCase jefeproyecto.CreateJefeProyecto) *CreateJefeProyectoController {
	return &CreateJefeProyectoController{useCase: useCase}
}

func (cjp_c *CreateJefeProyectoController) Execute(jefeproyecto entities.JefeProyecto) error{
	return cjp_c.useCase.Execute(jefeproyecto)
}