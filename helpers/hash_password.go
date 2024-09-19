package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hash)
}
