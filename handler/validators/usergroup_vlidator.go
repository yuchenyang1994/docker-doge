package validators

import (
	"docker-doge/db"
	"reflect"

	validator "gopkg.in/go-playground/validator.v8"
)

type UserGroupVlidator struct {
	GroupName string `json:"groupName" binding:"required,hasGroupName"`
}

func hasGroupName(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value,
	fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	d := db.GetDbInstance()
	if groupName, ok := field.Interface().(string); ok {
		usergroup := db.UserGroup{}
		if notFound := d.First(&usergroup, "group_name = ?", groupName).RecordNotFound(); notFound != true {
			return true
		}
	}
	return false
}

type UserGroupIdVlidator struct {
	GroupId uint `json:"groupId" binding:"required,hasGroupById"`
}

func hasGroupById(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value,
	fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	d := db.GetDbInstance()
	if groupId, ok := field.Interface().(uint); ok {
		usergroup := db.UserGroup{}
		if notFound := d.First(&usergroup, groupId).RecordNotFound(); notFound {
			return notFound
		}
	}
	return false
}
