package trabajador

import (
	"demo/src/server/domain"
	"demo/src/server/domain/entities"
)

type UpdateTrabajador struct {
	db domain.ITrabajador
}


func NewUpdateTrabajador(db domain.ITrabajador) *UpdateTrabajador {
	return &UpdateTrabajador{db: db}
}


func (ut *UpdateTrabajador) Execute(trabajador entities.Trabajador) error {
	return ut.db.Update(trabajador)
}