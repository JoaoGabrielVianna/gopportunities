package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
