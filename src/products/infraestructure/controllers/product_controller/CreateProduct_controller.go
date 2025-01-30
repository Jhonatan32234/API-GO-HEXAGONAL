package product_controller

import (
	"demo/src/products/application/product"
	"demo/src/products/domain/entities"
)


type CreateProductController struct{
	useCase product.CreateProduct
}

func NewCreateProductController(useCase product.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp_c *CreateProductController) Execute(product entities.Product) error{
	return cp_c.useCase.Execute(product)
}