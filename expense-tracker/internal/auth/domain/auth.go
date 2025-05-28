package auth_domain

type AuthService interface {
	VerifyPassword(password string, hash string) error
	GenerateToken(userId string) (string, error)
}