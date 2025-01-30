package trabajador

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
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