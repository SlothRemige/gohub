package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

func RegisterApiRoutes(router *gin.Engine) {
	v1Router := router.Group("api/v1")
	{
		authRouter := v1Router.Group("/auth")
		{
			sc := new(auth.SignupController)
			authRouter.POST("signup/phone/exist", sc.IsPhoneExist)

			authRouter.POST("signup/email/exist", sc.IsEmailExist)
		}
	}
}
