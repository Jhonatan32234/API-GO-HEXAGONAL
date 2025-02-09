package routes_jefeproyecto

import (
	"demo/src/server/jefeproyecto/infraestructure/handlers_jefeproyecto"
	"github.com/gin-gonic/gin"
)




func JefeProyectoRoutes( router *gin.Engine){
	router.GET("/coordinador",handlers_jefeproyecto.ListJefeProyectosHandler)
	router.POST("/coordinador",handlers_jefeproyecto.CreateJefeProyectoHandler)
	router.PUT("/coordinador/:id",handlers_jefeproyecto.UpdateJefeProyectosHandler)
	router.DELETE("/coordinador/:id",handlers_jefeproyecto.DeleteJefeProyectoHandler)
	router.GET("/coordinador/experiencia/:id",handlers_jefeproyecto.GetExperienceHandler)
}




/*
func createJefeProyectoHandler(c *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	if err := c.ShouldBindJSON(&jefeproyectos); err != nil {
		fmt.Println("Error en el binding de JSON:", err)
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
////esta parte
	ps := models_jefeproyecto.NewMySQLJP()
	createjefeproyecto := useCase_jefeproyecto.NewCreateJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewCreateJefeProyectoController(*createjefeproyecto)
////
	if err := controller.Execute(jefeproyectos); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "JefeProyecto creado exitosamente",
		"product": jefeproyectos, 
	})
}*/
/*
func listJefeProyectosHandler(c *gin.Context) {
	ps := models_jefeproyecto.NewMySQLJP()
	getAllJefeProyectos := useCase_jefeproyecto.NewListJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewListJefeProyectoController(*getAllJefeProyectos)

	jefeproyectos, err := controller.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(200, converted)
}

*/
/*
func updateJefeProyectosHandler(c *gin.Context) {
	var jefeproyectos entities.JefeProyecto

	if err := c.ShouldBindJSON(&jefeproyectos); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return 
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	jefeproyectos.Idjefeproyecto = int32(id) 

	ps := models_jefeproyecto.NewMySQLJP()
	updatejefeproyecto := useCase_jefeproyecto.NewUpdateJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewUpdateJefeProyectoController(*updatejefeproyecto)

	if err := controller.Execute(jefeproyectos); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "JefeProyecto actualizado correctamente"})
}
*/

/*
func deleteJefeProyectoHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	ps := models_jefeproyecto.NewMySQLJP()
	deleteJefeProyecto := useCase_jefeproyecto.NewDeleteJefeProyecto(ps)
	controller := controllers_jefeproyecto.NewDeleteJefeProyectoController(*deleteJefeProyecto)

	if err := controller.Execute(int32(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "JefeProyecto eliminado correctamente"})
}*/
