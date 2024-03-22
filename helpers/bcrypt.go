package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	pass := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)
	return string(hash)
}

func ComparePass(h, p []byte) bool {

	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
