package validators

import (
	"github.com/gin-gonic/gin/binding"
)

func RegisterV() {
	binding.Validator.RegisterValidation("RePassword", rePassword)
	binding.Validator.RegisterValidation("hasGroupName", hasGroupName)
	binding.Validator.RegisterValidation("hasGroupById", hasGroupById)
}
