package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (router *Router) UserRoutes(group *gin.RouterGroup) {
	group.POST("/create", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
	group.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
}
