package token_issuer

type TokenIssuer interface {
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	Parse(token string) (map[string]any, error)
}