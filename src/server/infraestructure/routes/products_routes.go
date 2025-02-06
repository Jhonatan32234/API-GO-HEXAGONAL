package routes

import (
	"demo/src/server/application/product"
	"demo/src/server/domain/entities"
	"demo/src/server/infraestructure/controllers/product_controller"
	"demo/src/server/infraestructure/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
GET /products ---> handler
POST /products --> handler (CreateProduct_controller)
PUT /products ---> handler
DELETE /products-> handler3
*/



func ProductRoutes( router *gin.Engine){
	router.GET("/products",listProductsHandler)
	router.POST("/products",createProductHandler)
	router.PUT("/products/:id",updateProductHandler)
	router.DELETE("/products/:id",deleteProductHandler)
}





func createProductHandler(c *gin.Context) {
	var products entities.Product

	// Decodificar el JSON recibido en la estructura product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	ps := models.NewMySQLP()
	createProduct := product.NewCreateProduct(ps)
	controller := product_controller.NewCreateProductController(*createProduct)

	// Llamar al controlador con los datos del producto
	if err := controller.Execute(products); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Responder con el producto incluyendo el ID generado
	c.JSON(201, gin.H{
		"message": "Producto creado exitosamente",
		"product": products, // Ahora incluirá el ID correcto
	})
}

func listProductsHandler(c *gin.Context) {
	ps :=  models.NewMySQLP()
	getAllProduct := product.NewListProduct(ps)
    controller := product_controller.NewListProductController(*getAllProduct)
    products, err := controller.Execute()
    // Add response handling
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, products)
}


func updateProductHandler(c *gin.Context) {
	var products entities.Product

	// Obtener el ID desde la URL
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	// Convertir el ID de la URL a entero
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	products.Id = int32(id) // Asignar el ID al producto recibido

	ps := models.NewMySQLP()
	updateProduct := product.NewUpdateProduct(ps)
	controller := product_controller.NewUpdateProductController(*updateProduct)

	if err := controller.Execute(products); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Producto actualizado correctamente"})
}

func deleteProductHandler(c *gin.Context) {
	// Obtener el ID desde la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	ps := models.NewMySQLP()
	deleteProduct := product.NewDeleteProduct(ps)
	controller := product_controller.NewDeleteProductController(*deleteProduct)

	if err := controller.Execute(int32(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Producto eliminado correctamente"})
}
