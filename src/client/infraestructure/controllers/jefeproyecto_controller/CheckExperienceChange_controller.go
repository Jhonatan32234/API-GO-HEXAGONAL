package jefeproyecto_controller

import (
	"demo/src/client/application/jefeproyecto"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckExperienceChangeController struct {
	useCase jefeproyecto.CheckExperienceChange
}

func NewCheckExperienceChangeController(useCase jefeproyecto.CheckExperienceChange) *CheckExperienceChangeController {
	return &CheckExperienceChangeController{useCase: useCase}
}

func (cec_c *CheckExperienceChangeController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	jefeID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	previousExperience, err := cec_c.useCase.Execute(int32(jefeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	timeout := time.After(30 * time.Second) 

	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case <-timeout:
				c.JSON(http.StatusNoContent, nil)
				return false
			default:
				currentExperience, err := cec_c.useCase.Execute(int32(jefeID))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return false
				}

				if currentExperience != previousExperience {
					previousExperience = currentExperience
					c.Writer.Write([]byte(fmt.Sprintf("Años de experiencia actualizado: %d\n", currentExperience)))
					c.Writer.Flush() 
					return false    
				}

				time.Sleep(1 * time.Second)
			}
		}
	})
}
