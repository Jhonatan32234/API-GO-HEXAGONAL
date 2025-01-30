package product

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(product entities.Product) error{
	return cp.db.Save(product)
}