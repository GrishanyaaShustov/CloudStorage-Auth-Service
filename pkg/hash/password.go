package hash

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int
}

func New(cost int) *Hasher {
	return &Hasher{cost: cost}
}

// HashPassword hashing password for saving in DB
func (h *Hasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPassword Compares raw password with hash from DB
func (h *Hasher) CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
