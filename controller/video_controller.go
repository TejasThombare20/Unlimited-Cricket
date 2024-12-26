package controller

import (
	"TejasThombare20/fampay/model"
	"TejasThombare20/fampay/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	service *service.YoutubeService
}

func NewVideoController(service *service.YoutubeService) *VideoController {
	return &VideoController{service: service}
}

func (h *VideoController) GetVideos(c *gin.Context) {
	var query model.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	videos, err := h.service.GetData(c.Request.Context(), page, pageSize)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch videos", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":      page,
		"page_size": pageSize,
		"data":      videos,
	})

}
