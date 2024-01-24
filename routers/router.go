package routers

import (
	"embed"
	"net/http"

	"github.com/limitcool/palworld-admin/global"
	"github.com/limitcool/palworld-admin/internal/handlers"
	"github.com/limitcool/palworld-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var f embed.FS

// Load loads the middlewares, routes, handlers.
func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(middleware.Serve("/", middleware.EmbedFolder(f, "dist")))
	router.GET("/", func(ctx *gin.Context) {
		data, err := f.ReadFile("static/dist/index.html")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	router.Use(middleware.Cors())

	// v1 router
	apiV1 := router.Group("/api/v1")
	auth := apiV1.Use(middleware.AdminPasswordMiddleware(global.Config.AdminPassword))
	{
		auth.GET("/config", handlers.GetConfig)
		auth.POST("/config", handlers.UpdateConfig)
	}
	return router
}
