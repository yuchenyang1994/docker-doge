package validators

import (
	"regexp"
	"testing"
)

func TestReEmail(t *testing.T) {
	if reg, err := regexp.Compile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`); err == nil {
		if ok := reg.MatchString("test@gmail.com"); ok != true {
			t.Error("email match error")
		}
		if ok := reg.MatchString("xxxx"); ok {
			t.Error("email match error")
		}
	}
}

func TestPassWord(t *testing.T) {
	if reg, err := regexp.Compile(`^[a-zA-Z]\w{5,17}$`); err == nil {
		if ok := reg.MatchString("test123456"); ok != true {
			t.Error("password match error")
		}
		if ok := reg.MatchString("xxxx.fdsdf12c"); ok {
			t.Error("password match error")
		}
	}
}
