package useCase_trabajador

import "demo/src/server/trabajador/domain"

type GetPosition struct {
    db domain.ITrabajador
}

func NewGetPosition(db domain.ITrabajador) *GetPosition {
    return &GetPosition{db: db}
}

func (csu *GetPosition) Execute(trabajadorID int32) (string, error) {
    return csu.db.GetPosition(trabajadorID)
}