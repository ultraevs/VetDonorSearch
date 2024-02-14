package router

import (
	"app/internal/api/contorllers"
	"app/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func (router *Router) UserRoutes(group *gin.RouterGroup) {
	group.GET("/questionnaire", middleware.AuthMiddleware(), controller.UserQuestionnaire)
	group.POST("/create_questionnaire", middleware.AuthMiddleware(), controller.CreateQuestionnaire)
	group.GET("/profile", middleware.AuthMiddleware(), controller.Profile)
	group.GET("profile/:key", controller.GetProfile)
}
