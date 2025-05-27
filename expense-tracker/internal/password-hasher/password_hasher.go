package password_hasher

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) error
}