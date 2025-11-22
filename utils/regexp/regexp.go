package utils

import (
	"errors"

	"github.com/dlclark/regexp2"
)

var PasswordPattern = regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,20}$`, regexp2.None)

func Match(re *regexp2.Regexp, s string) (err error) {
	ok, err := re.MatchString(s)
	if !ok {
		return errors.New("密码无效")
	}
	if err != nil {
		return err
	}
	return nil
}
