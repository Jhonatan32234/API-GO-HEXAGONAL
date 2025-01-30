package controller_jefeproyecto

import (
	"demo/src/products/application/jefeproyecto"
	"demo/src/products/domain/entities"
)

type ListJefeProyectoController struct{
	useCase jefeproyecto.ListJefeProyecto
}

func NewListJefeProyectoController(useCase jefeproyecto.ListJefeProyecto) *ListJefeProyectoController {
	return &ListJefeProyectoController{useCase: useCase}
}

func (ljp_c *ListJefeProyectoController) Execute() ([]entities.JefeProyecto, error) {
	return ljp_c.useCase.Execute()
}