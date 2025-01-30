package routes

import (
	"demo/src/products/application/jefeproyecto"
	"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/controllers/controller_jefeproyecto"
	"demo/src/products/infraestructure/models"
	"strconv"

	"github.com/gin-gonic/gin"
)




func JefeProyectoRoutes( router *gin.Engine){
	router.GET("/jefeproyecto",listJefeProyectosHandler)
	router.POST("/jefeproyecto",createJefeProyectoHandler)
	router.PUT("/jefeproyecto/:id",updateJefeProyectosHandler)
	router.DELETE("/jefeproyecto/:id",deleteJefeProyectoHandler)
}





func createJefeProyectoHandler(c *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	// Decodificar el JSON recibido en la estructura product
	if err := c.ShouldBindJSON(&jefeproyectos); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	ps := models.NewMySQLJP()
	createjefeproyecto := jefeproyecto.NewCreateJefeProyecto(ps)
	controller := controller_jefeproyecto.NewCreateJefeProyectoController(*createjefeproyecto)

	// Llamar al controlador con los datos del producto
	if err := controller.Execute(jefeproyectos); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Responder con el producto incluyendo el ID generado
	c.JSON(201, gin.H{
		"message": "JefeProyecto creado exitosamente",
		"product": jefeproyectos, // Ahora incluirá el ID correcto
	})
}

func listJefeProyectosHandler(c *gin.Context) {
	ps :=  models.NewMySQLJP()
	getAllJefeProyectos := jefeproyecto.NewListJefeProyecto(ps)
    controller := controller_jefeproyecto.NewListJefeProyectoController(*getAllJefeProyectos)
    jefeproyectos, err := controller.Execute()
    // Add response handling
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, jefeproyectos)
}


func updateJefeProyectosHandler(c *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	// Obtener el ID desde la URL
	if err := c.ShouldBindJSON(&jefeproyectos); err != nil {
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

	jefeproyectos.Idjefeproyecto = int32(id) // Asignar el ID al producto recibido

	ps := models.NewMySQLJP()
	updatejefeproyecto := jefeproyecto.NewUpdateJefeProyecto(ps)
	controller := controller_jefeproyecto.NewUpdateJefeProyectoController(*updatejefeproyecto)

	if err := controller.Execute(jefeproyectos); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "JefeProyecto actualizado correctamente"})
}

func deleteJefeProyectoHandler(c *gin.Context) {
	// Obtener el ID desde la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	ps := models.NewMySQLJP()
	deleteJefeProyecto := jefeproyecto.NewDeleteJefeProyecto(ps)
	controller := controller_jefeproyecto.NewDeleteJefeProyectoController(*deleteJefeProyecto)

	if err := controller.Execute(int32(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "JefeProyecto eliminado correctamente"})
}
