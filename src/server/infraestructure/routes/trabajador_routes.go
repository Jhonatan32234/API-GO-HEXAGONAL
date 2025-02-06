package routes

import (
	"demo/src/server/application/trabajador"
	"demo/src/server/domain/entities"
	"demo/src/server/infraestructure/controllers/controller_trabajador"
	"demo/src/server/infraestructure/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)




func TrabajadorRoutes( router *gin.Engine){
	router.GET("/trabajador",listTrabajadoresHandler)
	router.POST("/trabajador",createTrabajadorHandler)
	router.PUT("/trabajador/:id",updateTrabajadorHandler)
	router.DELETE("/trabajador/:id",deleteTrabajadorHandler)
	router.GET("/trabajador/salario/:id", GetSalaryHandler)
	router.GET("/trabajador/posicion/:id", GetPositionHandler)
}


func GetSalaryHandler(c *gin.Context) {
    ps := models.NewMySQLT()
    useCase := trabajador.NewGetSalary(ps)
    controller := controller_trabajador.NewGetSalaryController(*useCase)
    controller.Execute(c)
}

func GetPositionHandler(c *gin.Context) {
    ps := models.NewMySQLT()
    useCase := trabajador.NewGetPosition(ps)
    controller := controller_trabajador.NewGetPositionController(*useCase)
    controller.Execute(c)
}

func createTrabajadorHandler(c *gin.Context) {
	var trabajadores entities.Trabajador

	
	// Decodificar el JSON recibido en la estructura product
	if err := c.ShouldBindJSON(&trabajadores); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		c.JSON(400, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
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
	// Obtiene la conexión a la base de datos
	ps := models.NewMySQLT()
	getAllTrabajadores := trabajador.NewListTrabajador(ps)
	controller := controller_trabajador.NewListTrabajadorController(*getAllTrabajadores)

	// Ejecuta el controlador para obtener la lista de trabajadores
	trabajadores, err := controller.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Convierte los trabajadores en un formato con claves en minúsculas
	converted := make([]map[string]interface{}, 0)

	for _, t := range trabajadores {
		convertedTrabajador := map[string]interface{}{
			"idtrabajador":     t.Idtrabajador,
			"nombretrabajador": t.Nombretrabajador,
			"posicion":         t.Posicion,
			"telefono":         t.Telefono,
			"correo":           t.Correo,
			"salario":          t.Salario,
			"aniosexperiencia": t.Aniosexperiencia,
		}
		converted = append(converted, convertedTrabajador)
	}

	// Envía la respuesta con claves en minúsculas
	c.JSON(200, converted)
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


