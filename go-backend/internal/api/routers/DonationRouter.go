package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) DonationRouters(group *gin.RouterGroup) {
	group.POST("/check_donation", controller.CheckDonation)
	group.GET("/get_all_donations", controller.GetAllNewDonations)
	group.POST("/delete_application", controller.DeleteDonationApplication)
}
