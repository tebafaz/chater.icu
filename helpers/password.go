package helpers

import "golang.org/x/crypto/bcrypt"

func CreateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(bytes), err
}

func ComparePasswords(providedPass string, DBPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(DBPass), []byte(providedPass))
	return err == nil
}
