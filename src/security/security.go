package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(passwordDatabase, attemptPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordDatabase), []byte(attemptPassword))
}