package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing opennings")
		return
	}

	sendSuccess(c, "list-opening", openings)
}
