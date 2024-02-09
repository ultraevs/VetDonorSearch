package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) UserRoutes(group *gin.RouterGroup) {
	group.POST("/user_create", controller.UserCreate)
	group.POST("/clinic_create", controller.ClinicCreate)
	group.POST("/login", controller.Login)
}
