package bcrypt

import "golang.org/x/crypto/bcrypt"

// GetPasswordHash 负责对密码进行bcrypt加盐加密
func GetPasswordHash(password string) (string, error) {
	// bcrypt 自动生成盐，不需要手动处理
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
