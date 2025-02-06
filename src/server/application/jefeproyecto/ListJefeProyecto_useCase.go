package jefeproyecto

import (
	"demo/src/server/domain"
	"demo/src/server/domain/entities"
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