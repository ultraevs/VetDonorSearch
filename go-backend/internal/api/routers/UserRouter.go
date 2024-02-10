package router

import (
	"app/internal/api/contorllers"
	"app/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func (router *Router) UserRoutes(group *gin.RouterGroup) {
	group.GET("/profile", middleware.AuthMiddleware(), controller.UserProfile)
}
