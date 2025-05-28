package tokenissuer

type TokenIssuer interface {
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	GenerateRandomToken(size int) (string, error)
	Parse(token string) (map[string]any, error)
}