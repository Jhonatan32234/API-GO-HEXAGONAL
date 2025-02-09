package useCase_jefeproyecto

import (
	"demo/src/server/jefeproyecto/domain"
	"demo/src/server/jefeproyecto/domain/entities"
)

type UpdateJefeProyecto struct {
	db domain.IJefeProyecto
}


func NewUpdateJefeProyecto(db domain.IJefeProyecto) *UpdateJefeProyecto {
	return &UpdateJefeProyecto{db: db}
}


func (ujp *UpdateJefeProyecto) Execute(jefeproyecto entities.JefeProyecto) error {
	return ujp.db.Update(jefeproyecto)
}