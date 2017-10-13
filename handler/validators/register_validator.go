package validators

import (
	"github.com/gin-gonic/gin/binding"
)

func RegisterV() {
	binding.Validator.RegisterValidation("RePassword", rePassword)
}
