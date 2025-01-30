package jefeproyecto

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type CreateJefeProyecto struct {
	db domain.IJefeProyecto
}

func NewCreateJefeProyecto(db domain.IJefeProyecto) *CreateJefeProyecto {
	return &CreateJefeProyecto{db: db}
}

func (cjp *CreateJefeProyecto) Execute(jefeproyecto entities.JefeProyecto) error{
	return cjp.db.Save(jefeproyecto)
}