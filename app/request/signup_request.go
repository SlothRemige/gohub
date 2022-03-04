package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidatePhoneIsExist(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validator(data, rules, messages)
}

type EmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateEmailIsExist(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validator(data, rules, messages)
}
