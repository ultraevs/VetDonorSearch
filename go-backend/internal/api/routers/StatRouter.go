package router

import (
	"app/internal/api/contorllers"
	"github.com/gin-gonic/gin"
)

func (router *Router) StatRouters(group *gin.RouterGroup) {
	group.GET("/get_user_stats/:key", controller.GetUserStat)
	group.PUT("/update_user_stats", controller.UpdateUserStat)
	group.GET("/top_list", controller.TopDonationList)
}
