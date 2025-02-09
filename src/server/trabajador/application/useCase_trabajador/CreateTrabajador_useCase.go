package useCase_trabajador

import (
	"demo/src/server/trabajador/domain" 
	"demo/src/server/trabajador/domain/entities"
)

type CreateTrabajador struct {
	db domain.ITrabajador
}

func NewCreateTrabajador(db domain.ITrabajador) *CreateTrabajador {
	return &CreateTrabajador{db: db}
}

func (ct *CreateTrabajador) Execute(trabajador entities.Trabajador) error{
	return ct.db.Save(trabajador)
}