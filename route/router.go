package route

import (
	"TejasThombare20/fampay/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(routes *gin.Engine, ytController *controller.VideoController) {

	v1 := routes.Group("/api/v1")

	v1.GET("/lists", ytController.GetVideos)

}
