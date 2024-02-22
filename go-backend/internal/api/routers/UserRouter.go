package router

import (
	"app/internal/api/contorllers"
	"app/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func (router *Router) UserRoutes(group *gin.RouterGroup) {
	group.POST("/create_questionnaire", middleware.AuthMiddleware(), controller.CreateQuestionnaire)
	group.POST("/get_questionnaire", controller.GetQuestionnaire)
	group.GET("/profile", middleware.AuthMiddleware(), controller.Profile)
	group.GET("profile/:key", controller.GetProfile)
	group.POST("/mark", controller.MarkAsNeed)
	group.POST("create_other_info", controller.CreateUserOtherInfo)
	group.GET("other_info", controller.GetUserOtherInfo)
}
