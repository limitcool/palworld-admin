package routers

import (
	"github.com/limitcool/palworld-admin/global"
	"github.com/limitcool/palworld-admin/internal/handlers"
	"github.com/limitcool/palworld-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	g.Use(middleware.Cors())
	// v1 router
	apiV1 := g.Group("/api/v1")
	auth := apiV1.Use(middleware.AdminPasswordMiddleware(global.Config.AdminPassword))
	{
		auth.GET("/config", handlers.GetConfig)
	}
	return g
}
