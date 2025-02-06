package product

import (
	"demo/src/server/domain"
	"demo/src/server/domain/entities"
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