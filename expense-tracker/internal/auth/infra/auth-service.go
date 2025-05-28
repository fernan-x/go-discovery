package auth_infra

import (
	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	password_hasher "github.com/fernan-x/expense-tracker/internal/shared/password-hasher"
	token_issuer "github.com/fernan-x/expense-tracker/internal/shared/token-issuer"
)

type AuthServiceImpl struct {
	passwordHasher password_hasher.PasswordHasher
	tokenIssuer token_issuer.TokenIssuer
}

var _ auth_domain.AuthService = (*AuthServiceImpl)(nil)

func NewAuthService(passwordHasher password_hasher.PasswordHasher, tokenIssuer token_issuer.TokenIssuer) *AuthServiceImpl {
	return &AuthServiceImpl{passwordHasher, tokenIssuer}
}

func (s *AuthServiceImpl) VerifyPassword(password string, hash string) error {
	return s.passwordHasher.Verify(password, hash)
}

func (s *AuthServiceImpl) GenerateToken(userId string) (string, error) {
	return s.tokenIssuer.Generate(userId)
}