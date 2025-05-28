package token_issuer

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenIssuer struct {
	secret []byte
}

var _ TokenIssuer = (*JwtTokenIssuer)(nil)

func NewJwtTokenIssuer(secret []byte) *JwtTokenIssuer {
	return &JwtTokenIssuer{secret}
}

func generateToken(claims jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *JwtTokenIssuer) GenerateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 1 hours
		"iss": "expense-tracker",
		"type": "access",
	}

	return generateToken(claims, t.secret)
}

func (t *JwtTokenIssuer) GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
		"iss": "expense-tracker",
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *JwtTokenIssuer) Parse(token string) (map[string]any, error) {
	jwtDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})
	if err != nil {
		return nil, err
	}

	return jwtDecoded.Claims.(jwt.MapClaims), nil
}