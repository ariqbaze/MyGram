package helpers

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
