package jefeproyecto

import (
	"demo/src/server/domain"
	"demo/src/server/domain/entities"
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