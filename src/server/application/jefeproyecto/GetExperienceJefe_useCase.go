package jefeproyecto

import (
	"demo/src/server/domain"
)

type GetExperience struct {
    db domain.IJefeProyecto
}

func NewGetExperience(db domain.IJefeProyecto) *GetExperience {
    return &GetExperience{db: db}
}

func (ceu *GetExperience) Execute(jefeproyectoID int32) (int32, error) {
    return ceu.db.GetExperience(jefeproyectoID)
}