package controllers_trabajador

import (
	"demo/src/server/trabajador/application/useCase_trabajador"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetSalaryController struct {
    useCase useCase_trabajador.GetSalary
}

func NewGetSalaryController(useCase useCase_trabajador.GetSalary) *GetSalaryController {
    return &GetSalaryController{useCase: useCase}
}

func (csu_c *GetSalaryController) Execute(c *gin.Context) {
    idParam := c.Param("id")
    trabajadorID, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    salary, err := csu_c.useCase.Execute(int32(trabajadorID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    log.Println("Salario: ", salary)

    c.JSON(http.StatusOK, gin.H{"salario": salary})
}