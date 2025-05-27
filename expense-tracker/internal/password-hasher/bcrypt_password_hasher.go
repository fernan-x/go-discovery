package password_hasher

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {}

var _ PasswordHasher = (*BcryptPasswordHasher)(nil)

func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (h *BcryptPasswordHasher) Verify(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}