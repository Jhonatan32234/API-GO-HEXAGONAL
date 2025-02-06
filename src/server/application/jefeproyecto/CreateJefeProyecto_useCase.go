package jefeproyecto

import (
	"demo/src/server/domain"
	"demo/src/server/domain/entities"
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