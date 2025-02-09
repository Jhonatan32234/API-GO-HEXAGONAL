package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetExperienceController struct {
    useCase useCase_jefeproyecto.GetExperience
}

func NewGetExperienceController(useCase useCase_jefeproyecto.GetExperience) *GetExperienceController {
    return &GetExperienceController{useCase: useCase}
}

func (ceu_c *GetExperienceController) Execute(c *gin.Context) {
    idParam := c.Param("id")
    jefeproyectoID, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    experiencia, err := ceu_c.useCase.Execute(int32(jefeproyectoID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"anosexperiencia": experiencia})
}

/*
package controllers_jefeproyecto

import (
	"demo/src/server/jefeproyecto/application/useCase_jefeproyecto"
	"demo/src/server/jefeproyecto/infraestructure/models_jefeproyecto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetExperienceController struct {
	useCase useCase_jefeproyecto.GetExperience
}

func NewGetExperienceController(useCase useCase_jefeproyecto.GetExperience) *GetExperienceController {
	return &GetExperienceController{useCase: useCase}
}

// Mueve la inicialización aquí
func GetExperienceHandler(c *gin.Context) {
	ps := models_jefeproyecto.NewMySQLJP()
	useCase := useCase_jefeproyecto.NewGetExperience(ps)
	controller := NewGetExperienceController(*useCase)
	controller.Execute(c)
}

// Mantén la lógica de ejecución aquí
func (ceu_c *GetExperienceController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	jefeproyectoID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	experiencia, err := ceu_c.useCase.Execute(int32(jefeproyectoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"anosexperiencia": experiencia})
}
*/