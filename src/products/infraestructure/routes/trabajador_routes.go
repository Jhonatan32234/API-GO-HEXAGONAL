package routes

import (
	"demo/src/products/application/trabajador"
	"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/controllers/controller_trabajador"
	"demo/src/products/infraestructure/models"
	"strconv"

	"github.com/gin-gonic/gin"
)




func TrabajadorRoutes( router *gin.Engine){
	router.GET("/trabajador",listTrabajadoresHandler)
	router.POST("/trabajador",createTrabajadorHandler)
	router.PUT("/trabajador/:id",updateTrabajadorHandler)
	router.DELETE("/trabajador/:id",deleteTrabajadorHandler)
}





func createTrabajadorHandler(c *gin.Context) {
	var trabajadores entities.Trabajador

	// Decodificar el JSON recibido en la estructura product
	if err := c.ShouldBindJSON(&trabajadores); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	ps := models.NewMySQLT()
	createTrabajador := trabajador.NewCreateTrabajador(ps)
	controller := controller_trabajador.NewCreateTrabajadorController(*createTrabajador)

	// Llamar al controlador con los datos del producto
	if err := controller.Execute(trabajadores); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Responder con el producto incluyendo el ID generado
	c.JSON(201, gin.H{
		"message": "Trabajador creado exitosamente",
		"product": trabajadores, // Ahora incluirá el ID correcto
	})
}

func listTrabajadoresHandler(c *gin.Context) {
	ps :=  models.NewMySQLT()
	getAllTrabajadores := trabajador.NewListTrabajador(ps)
    controller := controller_trabajador.NewListTrabajadorController(*getAllTrabajadores)
    trabajadores, err := controller.Execute()
    // Add response handling
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, trabajadores)
}


func updateTrabajadorHandler(c *gin.Context) {
	var trabajadores entities.Trabajador

	// Obtener el ID desde la URL
	if err := c.ShouldBindJSON(&trabajadores); err != nil {
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

	trabajadores.Idtrabajador = int32(id) // Asignar el ID al producto recibido

	ps := models.NewMySQLT()
	updateTrabajador := trabajador.NewUpdateTrabajador(ps)
	controller := controller_trabajador.NewUpdateTrabajadorController(*updateTrabajador)

	if err := controller.Execute(trabajadores); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Trabajador actualizado correctamente"})
}

func deleteTrabajadorHandler(c *gin.Context) {
	// Obtener el ID desde la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	ps := models.NewMySQLT()
	deleteTrabajador := trabajador.NewDeleteTrabajador(ps)
	controller := controller_trabajador.NewDeleteTrabajadorController(*deleteTrabajador)

	if err := controller.Execute(int32(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Trabajador eliminado correctamente"})
}
