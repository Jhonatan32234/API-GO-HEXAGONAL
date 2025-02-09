package useCase_jefeproyecto

import (
	"demo/src/server/jefeproyecto/domain"
	"demo/src/server/jefeproyecto/domain/entities"
)

 

type ListJefeProyecto struct {
	db domain.IJefeProyecto
}

func NewListJefeProyecto(db domain.IJefeProyecto) *ListJefeProyecto {
	return &ListJefeProyecto{db: db}
}

func (ljp *ListJefeProyecto) Execute() ([]entities.JefeProyecto, error){
	return ljp.db.GetAll()
}