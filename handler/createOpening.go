package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabrielvianna/gopportunities/schemas"
)

func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validating error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)
		sendError(c, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(c, "CreateOpeningHandler", opening)

}
