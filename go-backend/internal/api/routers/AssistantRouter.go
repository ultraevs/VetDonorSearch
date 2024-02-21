package router

import (
	controller "app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) AssistantRoutes(group *gin.RouterGroup) {
	group.POST("/message", controller.Message)
}
