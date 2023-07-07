package middleware

import (
	"github.com/Yasinqurni/be-project/src/app/user/http/request"
	"github.com/go-playground/validator/v10"
)

func UserValidateStruct(user request.UserRequest) bool {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return false
	}
	return true
}
