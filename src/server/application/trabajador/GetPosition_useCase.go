package trabajador

import (
	"demo/src/server/domain"
)

type GetPosition struct {
    db domain.ITrabajador
}

func NewGetPosition(db domain.ITrabajador) *GetPosition {
    return &GetPosition{db: db}
}

func (csu *GetPosition) Execute(trabajadorID int32) (string, error) {
    return csu.db.GetPosition(trabajadorID)
}