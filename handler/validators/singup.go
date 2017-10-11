package validators

import (
	"reflect"
	"regexp"

	"docker-doge/db"

	validator "gopkg.in/go-playground/validator.v8"
)

type SingupVlidator struct {
	Email       string `json:"email" binding:"required, reEmail"`
	Password    string `json:"password" binding:"required, rePassword"`
	UserGroupID uint   `json:"userGroupId" binding:"required, checkUserGroupId"`
}

// reEmail ...
func reEmail(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if email, ok := field.Interface().(string); ok {
		if reg, err := regexp.Compile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`); err == nil {
			ok := reg.MatchString(email)
			if ok == true {
				return true
			}
			return false
		}
	}
	return false
}

// rePassword ...
func rePassword(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if password, ok := field.Interface().(string); ok {
		if reg, err := regexp.Compile(`^[a-zA-Z]\w{5,17}$`); err == nil {
			if reg.MatchString(password) {
				return true
			}
			return false
		}
	}
	return false
}

func checkUserGroupId(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	d := db.GetDbInstance()
	defer d.Close()
	if userGroupId, ok := field.Interface().(uint); ok {
		user := &db.User{}
		if ok := d.First(user, userGroupId).RecordNotFound(); ok {
			return false
		}
		return true
	}
	return false

}
