package auth_domain

type AuthService interface {
	VerifyPassword(password string, hash string) error
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	ParseToken(token string) (map[string]any, error)
}
