package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) CardsRoutes(group *gin.RouterGroup) {
	group.GET("/get_user_cards", controller.GetUserCard)
	group.GET("/get_clinic_cards", controller.GetClinicCard)
}
