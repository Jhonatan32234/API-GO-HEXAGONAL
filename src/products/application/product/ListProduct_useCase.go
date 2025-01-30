package product

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

 

type ListProduct struct {
	db domain.IProduct
}

func NewListProduct(db domain.IProduct) *ListProduct {
	return &ListProduct{db: db}
}

func (lp *ListProduct) Execute() ([]entities.Product, error){
	return lp.db.GetAll()
}