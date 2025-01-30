package product_controller

import (
	"demo/src/products/application/product"
	"demo/src/products/domain/entities"
)

type ListProductController struct{
	useCase product.ListProduct
}

func NewListProductController(useCase product.ListProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (lp_c *ListProductController) Execute() ([]entities.Product, error) {
	return lp_c.useCase.Execute()
}