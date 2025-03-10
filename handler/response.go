package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/gopportunities/schemas"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorcode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorcode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
