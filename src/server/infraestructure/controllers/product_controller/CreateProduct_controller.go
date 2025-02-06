package product_controller

import (
	"demo/src/server/application/product"
	"demo/src/server/domain/entities"
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