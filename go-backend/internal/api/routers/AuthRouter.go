package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) AuthRoutes(group *gin.RouterGroup) {
	group.POST("/user_create", controller.UserCreate)
	group.POST("/clinic_create", controller.ClinicCreate)
	group.POST("/login", controller.Login)
	group.POST("/forgot", controller.ForgotPassword)
	group.GET("/newpass", controller.NewPassword)
	group.POST("/newpass", controller.PostNewPassword)
}
