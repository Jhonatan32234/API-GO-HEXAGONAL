package domain

import "demo/src/server/jefeproyecto/domain/entities"

type IJefeProyecto interface {
    GetAll() ([]entities.JefeProyecto, error)
	Save(jefeproyecto entities.JefeProyecto)error
	Update(jefeproyecto entities.JefeProyecto) error
	Delete(jefeproyectoID int32) error
	GetExperience(jefeProyectoID int32) (int32, error)
}


