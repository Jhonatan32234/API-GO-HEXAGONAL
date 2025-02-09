package domain

import "demo/src/server/trabajador/domain/entities"

type ITrabajador interface {
    GetAll() ([]entities.Trabajador, error)
	Save(trabajador entities.Trabajador)error
	Update(trabajador entities.Trabajador) error
	Delete(trabajadorID int32) error
	GetSalary(trabajadorID int32) (int32, error)
	GetPosition(trabajadorID int32) (string, error)
}


