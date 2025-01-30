package domain

import "demo/src/products/domain/entities"

type IJefeProyecto interface {
    GetAll() ([]entities.JefeProyecto, error)
	Save(jefeproyecto entities.JefeProyecto)error
	Update(jefeproyecto entities.JefeProyecto) error
	Delete(jefeproyectoID int32) error
}


