package routes

import (
	"demo/src/server/application/jefeproyecto"
	"demo/src/server/domain/entities"
	"demo/src/server/infraestructure/controllers/controller_jefeproyecto"
	"demo/src/server/infraestructure/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)




func JefeProyectoRoutes( router *gin.Engine){
	router.GET("/jefeproyecto",listJefeProyectosHandler)
	router.POST("/jefeproyecto",createJefeProyectoHandler)
	router.PUT("/jefeproyecto/:id",updateJefeProyectosHandler)
	router.DELETE("/jefeproyecto/:id",deleteJefeProyectoHandler)
	router.GET("/jefeproyecto/experiencia/:id",GetExperienceHandler)
}



func GetExperienceHandler(c *gin.Context) {
    ps := models.NewMySQLJP()
    useCase := jefeproyecto.NewGetExperience(ps)
    controller := controller_jefeproyecto.NewGetExperienceController(*useCase)
    controller.Execute(c)
}



func createJefeProyectoHandler(c *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	// Decodificar el JSON recibido en la estructura product
	if err := c.ShouldBindJSON(&jefeproyectos); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
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
	// Obtiene la conexión a la base de datos
	ps := models.NewMySQLJP()
	getAllJefeProyectos := jefeproyecto.NewListJefeProyecto(ps)
	controller := controller_jefeproyecto.NewListJefeProyectoController(*getAllJefeProyectos)

	// Ejecuta el controlador para obtener la lista de jefes de proyecto
	jefeproyectos, err := controller.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Convierte los jefes de proyecto en un formato con claves en minúsculas
	converted := make([]map[string]interface{}, 0)

	for _, jp := range jefeproyectos {
		convertedJefeProyecto := map[string]interface{}{
			"idjefeproyecto":   jp.Idjefeproyecto,
			"nombrejefe":       jp.Nombrejefe,
			"telefono":         jp.Telefono,
			"correo":           jp.Correo,
			"salario":          jp.Salario,
			"aniosexperiencia": jp.Aniosexperiencia,
		}
		converted = append(converted, convertedJefeProyecto)
	}

	// Envía la respuesta con claves en minúsculas
	c.JSON(200, converted)
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
