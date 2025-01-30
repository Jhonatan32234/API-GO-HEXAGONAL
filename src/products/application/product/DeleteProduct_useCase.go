package product

import "demo/src/products/domain"

type DeleteProduct struct {
	db domain.IProduct
}


func NewDeleteProduct(db domain.IProduct) *DeleteProduct {
	return &DeleteProduct{db:db}
}


func (dp *DeleteProduct) Execute(productID int32) error {
	return dp.db.Delete(productID)
}