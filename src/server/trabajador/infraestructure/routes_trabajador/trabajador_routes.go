package routes_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"demo/src/server/trabajador/infraestructure/controllers_trabajador"
	"demo/src/server/trabajador/infraestructure/handlers_trabajador"
	"demo/src/server/trabajador/infraestructure/models_trabajador"

	"github.com/gin-gonic/gin"
)




func TrabajadorRoutes( router *gin.Engine){
	router.GET("/trabajador",handlers_trabajador.ListTrabajadoresHandler)
	router.POST("/trabajador",handlers_trabajador.CreateTrabajadorHandler)
	router.PUT("/trabajador/:id",handlers_trabajador.UpdateTrabajadorHandler)
	router.DELETE("/trabajador/:id",handlers_trabajador.DeleteTrabajadorHandler)
	router.GET("/trabajador/salario/:id", GetSalaryHandler)
	router.GET("/trabajador/posicion/:id", handlers_trabajador.GetPositionHandler)
}


func GetSalaryHandler(c *gin.Context) {
    ps := models_trabajador.NewMySQLT()
    useCase := useCase_trabajador.NewGetSalary(ps)
    controller := controllers_trabajador.NewGetSalaryController(*useCase)
    controller.Execute(c)
}

/*

func GetPositionHandler(c *gin.Context) {
    ps := models_trabajador.NewMySQLT()
    useCase := useCase_trabajador.NewGetPosition(ps)
    controller := controllers_trabajador.NewGetPositionController(*useCase)
    controller.Execute(c)
}
/*
func createTrabajadorHandler(c *gin.Context) {
	var trabajadores entities.Trabajador

	
	if err := c.ShouldBindJSON(&trabajadores); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		c.JSON(400, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	ps := models_trabajador.NewMySQLT()
	createTrabajador := useCase_trabajador.NewCreateTrabajador(ps)
	controller := controllers_trabajador.NewCreateTrabajadorController(*createTrabajador)

	if err := controller.Execute(trabajadores); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "Trabajador creado exitosamente",
		"product": trabajadores, 
	})
}
*/

/*

func listTrabajadoresHandler(c *gin.Context) {
	ps := models_trabajador.NewMySQLT()
	getAllTrabajadores := useCase_trabajador.NewListTrabajador(ps)
	controller := controllers_trabajador.NewListTrabajadorController(*getAllTrabajadores)

	trabajadores, err := controller.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(200, converted)
}

*//*

func updateTrabajadorHandler(c *gin.Context) {
	var trabajadores entities.Trabajador

	if err := c.ShouldBindJSON(&trabajadores); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	trabajadores.Idtrabajador = int32(id)

	ps := models_trabajador.NewMySQLT()
	updateTrabajador := useCase_trabajador.NewUpdateTrabajador(ps)
	controller := controllers_trabajador.NewUpdateTrabajadorController(*updateTrabajador)

	if err := controller.Execute(trabajadores); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Trabajador actualizado correctamente"})
}
*/


/*
func deleteTrabajadorHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	ps := models_trabajador.NewMySQLT()
	deleteTrabajador := useCase_trabajador.NewDeleteTrabajador(ps)
	controller := controllers_trabajador.NewDeleteTrabajadorController(*deleteTrabajador)

	if err := controller.Execute(int32(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Trabajador eliminado correctamente"})
}
*/

