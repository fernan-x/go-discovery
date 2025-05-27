package auth_domain

type AuthService interface {
	VerifyPassword(password string, hash string) error
}