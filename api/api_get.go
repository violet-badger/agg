package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPerson(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
