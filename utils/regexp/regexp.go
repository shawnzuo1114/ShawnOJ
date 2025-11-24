package utils

import (
	"errors"

	"github.com/dlclark/regexp2"
)

var PasswordPattern = regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,20}$`, regexp2.None)
var EmailPattern = regexp2.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, regexp2.None)

func Match(re *regexp2.Regexp, s string) (err error) {
	ok, err := re.MatchString(s)
	if !ok {
		if s == PasswordPattern.String() {
			return errors.New("密码无效")
		} else {
			return errors.New("邮箱无效")
		}
	}
	if err != nil {
		return err
	}
	return nil
}
