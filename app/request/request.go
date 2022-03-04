package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func validator(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,     // request object
		Rules:         rules,    // rules map
		Messages:      messages, // custom message map (Optional)
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}

func ValidateStruct(data interface{}, ctx *gin.Context, handleValidator ValidatorFunc) bool {

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"data":    err.Error(),
		})
		fmt.Println("StatusUnprocessableEntity:" + err.Error())
		return false
	}

	fmt.Println("request", data)
	errs := handleValidator(data, ctx)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求验证不通过，具体请查看 errors",
			"data":    errs,
		})
		return false
	}

	return true
}
