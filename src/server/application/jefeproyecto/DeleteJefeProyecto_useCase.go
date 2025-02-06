package jefeproyecto

import "demo/src/server/domain"

type DeleteJefeProyecto struct {
	db domain.IJefeProyecto
}


func NewDeleteJefeProyecto(db domain.IJefeProyecto) *DeleteJefeProyecto {
	return &DeleteJefeProyecto{db:db}
}


func (djp *DeleteJefeProyecto) Execute(jefeproyectoID int32) error {
	return djp.db.Delete(jefeproyectoID)
}