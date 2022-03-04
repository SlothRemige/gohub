package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	requests "gohub/app/request"
	"gohub/models/user"
	"net/http"
)

type SignupController struct {
	v1.BaseController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	request := requests.PhoneExistRequest{}
	if validatorRes := requests.ValidateStruct(&request, ctx, requests.ValidatePhoneIsExist); !validatorRes {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"data":    user.IsExistPhone(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(ctx *gin.Context) {

	request := requests.EmailExistRequest{}
	if validatorRes := requests.ValidateStruct(&request, ctx, requests.ValidateEmailIsExist); !validatorRes {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"data":    user.IsExistEmail(request.Email),
	})
}
