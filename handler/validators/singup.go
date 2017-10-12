package validators

import (
	"reflect"

	"regexp"

	validator "gopkg.in/go-playground/validator.v8"
)

type SingupVlidator struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,RePassword"`
	UserGroupID uint   `json:"userGroupId" binding:"required"`
}

func rePassword(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if password, ok := field.Interface().(string); ok {
		if matched, err := regexp.MatchString(`^[a-zA-Z]\w{5,17}$`, password); err == nil {
			return matched
		}

	}
	return false

}
