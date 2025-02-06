package controller_jefeproyecto

import (
	"demo/src/server/application/jefeproyecto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetExperienceController struct {
    useCase jefeproyecto.GetExperience
}

func NewGetExperienceController(useCase jefeproyecto.GetExperience) *GetExperienceController {
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