package controller_trabajador

import (
	"demo/src/server/application/trabajador"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPositionController struct {
    useCase trabajador.GetPosition
}

func NewGetPositionController(useCase trabajador.GetPosition) *GetPositionController {
    return &GetPositionController{useCase: useCase}
}

func (csu_c *GetPositionController) Execute(c *gin.Context) {
    idParam := c.Param("id")
    trabajadorID, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    posicion, err := csu_c.useCase.Execute(int32(trabajadorID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"posicion": posicion})
}