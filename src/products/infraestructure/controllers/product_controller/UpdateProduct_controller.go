package product_controller

import (
	"demo/src/products/application/product"
	"demo/src/products/domain/entities"
)
type UpdateProductController struct {
	useCase product.UpdateProduct
}

func NewUpdateProductController(useCase product.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCase: useCase}
}

func (up_c *UpdateProductController) Execute(product entities.Product) error {
	return up_c.useCase.Execute(product)
}
