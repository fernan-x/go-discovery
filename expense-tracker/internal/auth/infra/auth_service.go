package authinfra

import (
	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	tokenissuer "github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
)

type AuthServiceImpl struct {
	passwordHasher passwordhasher.PasswordHasher
	tokenIssuer tokenissuer.TokenIssuer
}

var _ authdomain.AuthService = (*AuthServiceImpl)(nil)

func NewAuthService(passwordHasher passwordhasher.PasswordHasher, tokenIssuer tokenissuer.TokenIssuer) *AuthServiceImpl {
	return &AuthServiceImpl{passwordHasher, tokenIssuer}
}

func (s *AuthServiceImpl) VerifyPassword(password string, hash string) error {
	return s.passwordHasher.Verify(password, hash)
}

func (s *AuthServiceImpl) GenerateAccessToken(userId string) (string, error) {
	return s.tokenIssuer.GenerateAccessToken(userId)
}

func (s *AuthServiceImpl) GenerateRefreshToken(userId string) (string, error) {
	return s.tokenIssuer.GenerateRefreshToken(userId)
}

func (s *AuthServiceImpl) ParseToken(token string) (map[string]any, error) {
	return s.tokenIssuer.Parse(token)
}