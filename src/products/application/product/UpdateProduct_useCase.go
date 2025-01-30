package product

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type UpdateProduct struct {
	db domain.IProduct
}


func NewUpdateProduct(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}


func (up *UpdateProduct) Execute(product entities.Product) error {
	return up.db.Update(product)
}