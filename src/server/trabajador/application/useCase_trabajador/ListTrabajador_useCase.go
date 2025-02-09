package useCase_trabajador

import (
	"demo/src/server/trabajador/domain"
	"demo/src/server/trabajador/domain/entities"
)

 

type ListTrabajador struct {
	db domain.ITrabajador
}

func NewListTrabajador(db domain.ITrabajador) *ListTrabajador {
	return &ListTrabajador{db: db}
}

func (lp *ListTrabajador) Execute() ([]entities.Trabajador, error){
	return lp.db.GetAll()
}