package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetUpRouter(router *gin.Engine) {

	RegisterGlobalMiddleRouter(router)

	routes.RegisterApiRoutes(router)

	SetUp4o4HandlerRouter(router)
}

func RegisterGlobalMiddleRouter(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func SetUp4o4HandlerRouter(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		acceptHeader := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptHeader, "text/html") {
			ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": http.StatusText(http.StatusNotFound),
				"data":    nil,
			})
		}
	})
}
