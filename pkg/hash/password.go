package hash

import "golang.org/x/crypto/bcrypt"

// HashPassword hashing password for saving in DB
func HashPassword(password string) (string, error) {
	const cost = bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPassword Compares raw password with hash from DB
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
