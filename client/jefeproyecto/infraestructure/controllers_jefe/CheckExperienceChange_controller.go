package controllers_jefe

import (
	//"demo/src/server/application/jefeproyecto"
    "demo/client/jefeproyecto/application/useCase_jefe"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckExperienceChangeController struct {
	useCase useCase_jefe.CheckExperienceChange
}

func NewCheckExperienceChangeController(useCase useCase_jefe.CheckExperienceChange) *CheckExperienceChangeController {
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
	//log.Println("Años de experiencia: ", previousExperience)
	c.SSEvent("experience_update", fmt.Sprintf("Años de experiencia actual: %d", previousExperience))
	c.Writer.Flush()

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case <-timeout:
				c.JSON(http.StatusNoContent, nil)
				return false
			case <-ticker.C:
				currentExperience, err := cec_c.useCase.Execute(int32(jefeID))
				if err != nil {
					// En lugar de cerrar la conexión, simplemente registra el error
					fmt.Fprintf(w, "event:error\ndata:%s\n\n", err.Error())
					c.Writer.Flush()
					return false
				}

				// Si hay un cambio en la experiencia, lo enviamos
				if currentExperience != previousExperience {
					previousExperience = currentExperience
					c.SSEvent("",fmt.Sprintf("Años de experiencia actualizado: %d", currentExperience))
					c.Writer.Flush()
				}
			}
		}
	})
}
