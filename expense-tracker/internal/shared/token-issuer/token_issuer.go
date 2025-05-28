package token_issuer

type TokenIssuer interface {
	Generate(userId string) (string, error)
	Parse(token string) (map[string]any, error)
}