package domain

import "demo/src/products/domain/entities"

type ITrabajador interface {
    GetAll() ([]entities.Trabajador, error)
	Save(trabajador entities.Trabajador)error
	Update(trabajador entities.Trabajador) error
	Delete(trabajadorID int32) error
}


