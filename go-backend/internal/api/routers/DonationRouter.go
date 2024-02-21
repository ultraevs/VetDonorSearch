package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) DonationRouters(group *gin.RouterGroup) {
	group.POST("/check_donation", controller.CheckDonation)
}
