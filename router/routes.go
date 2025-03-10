package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/joaogabrielvianna/gopportunities/docs"
	"github.com/joaogabrielvianna/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
<<<<<<< HEAD
=======
	"github.com/swaggo/swag/example/override/docs"
>>>>>>> 7b9bd8647c559d091685b7b3baff8af2182912ec
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
<<<<<<< HEAD
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
=======
	basePath := "/ap1/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group("/api/v1")
>>>>>>> 7b9bd8647c559d091685b7b3baff8af2182912ec
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

<<<<<<< HEAD
	// // Initialize swagger docs
=======
	// Initialize swagger docs
>>>>>>> 7b9bd8647c559d091685b7b3baff8af2182912ec
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
