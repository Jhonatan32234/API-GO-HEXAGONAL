package controllers_trabajador

import (
	"demo/client/trabajador/application/useCase_trabajador"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckSalaryChangeController struct {
	useCase useCase_trabajador.CheckSalaryChange
}

func NewCheckSalaryChangeController(useCase useCase_trabajador.CheckSalaryChange) *CheckSalaryChangeController {
	return &CheckSalaryChangeController{useCase: useCase}
}

func (csc_c *CheckSalaryChangeController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	trabajadorID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var previousSalary int32

	c.Stream(func(w io.Writer) bool {
		for {
			currentSalary, err := csc_c.useCase.Execute(int32(trabajadorID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
			log.Println("Salario: ", currentSalary)
			log.Println("Salario anterior: ", previousSalary)

			if currentSalary != previousSalary {
				log.Println("Salario actualizado")
				previousSalary = currentSalary
				c.Writer.Write([]byte(fmt.Sprintf("Salario actualizado: %d\n", currentSalary)))
				c.Writer.Flush() 
			}

			time.Sleep(5 * time.Second) 
		}
	})
}
