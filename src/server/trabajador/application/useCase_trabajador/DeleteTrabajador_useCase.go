package useCase_trabajador

import "demo/src/server/trabajador/domain"

type DeleteTrabajador struct {
	db domain.ITrabajador
}


func NewDeleteTrabajador(db domain.ITrabajador) *DeleteTrabajador {
	return &DeleteTrabajador{db:db}
}


func (dt *DeleteTrabajador) Execute(trabajdorID int32) error {
	return dt.db.Delete(trabajdorID)
}