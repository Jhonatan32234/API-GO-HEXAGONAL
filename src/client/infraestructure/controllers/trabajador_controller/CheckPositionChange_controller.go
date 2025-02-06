package trabajador_controller

import (
	"demo/src/client/application/trabajador"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckPositionChangeController struct {
	useCase trabajador.CheckPositionChange
}

func NewCheckPositionChangeController(useCase trabajador.CheckPositionChange) *CheckPositionChangeController {
	return &CheckPositionChangeController{useCase: useCase}
}

func (cpc_c *CheckPositionChangeController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	trabajadorID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var previousPosition string

	c.Stream(func(w io.Writer) bool {
		for {
			currentPosition, err := cpc_c.useCase.Execute(int32(trabajadorID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}

			if currentPosition != previousPosition {
				previousPosition = currentPosition
				c.Writer.Write([]byte(fmt.Sprintf("Posicion actualizado: %s\n", currentPosition)))
				c.Writer.Flush() 
			}

			time.Sleep(5 * time.Second) 
		}
	})
}
