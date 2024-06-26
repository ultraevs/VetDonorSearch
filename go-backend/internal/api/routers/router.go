package router

import (
	_ "app/docs"
	controller "app/internal/api/contorllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter() Router {
	return Router{engine: gin.Default()}
}

func (router *Router) Run(port string) error {
	router.Setup()
	return router.engine.Run(":" + port)
}

func (router *Router) Setup() {
	gin.SetMode(gin.DebugMode)
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://vetdonor.shmyaks.ru"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "*"},
		AllowCredentials: true,
	}))
	router.engine.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.engine.GET("v1/image/:key", controller.GetImage)
	v1 := router.engine.Group("/v1")
	router.AuthRoutes(v1)
	router.UserRoutes(v1)
	router.CardsRoutes(v1)
	router.StatRouters(v1)
	router.DonationRouters(v1)
	router.AssistantRoutes(v1)
}
