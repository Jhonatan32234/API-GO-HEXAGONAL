package useCase_trabajador

import "demo/src/server/trabajador/domain"

type GetSalary struct {
    db domain.ITrabajador
}

func NewGetSalary(db domain.ITrabajador) *GetSalary {
    return &GetSalary{db: db}
}

func (csu *GetSalary) Execute(trabajadorID int32) (int32, error) {
    return csu.db.GetSalary(trabajadorID)
}