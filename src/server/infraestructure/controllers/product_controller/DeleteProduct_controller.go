package product_controller

import (
	"demo/src/server/application/product"
)

type DeleteProductController struct {
	useCase product.DeleteProduct
}

func NewDeleteProductController(useCase product.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (dp_c *DeleteProductController) Execute(productID int32) error {
	return dp_c.useCase.Execute(productID)
}