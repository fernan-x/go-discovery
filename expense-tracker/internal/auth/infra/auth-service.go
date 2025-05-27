package auth_infra

import (
	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	password_hasher "github.com/fernan-x/expense-tracker/internal/password-hasher"
)

type AuthServiceImpl struct {
	passwordHasher password_hasher.PasswordHasher
}

var _ auth_domain.AuthService = (*AuthServiceImpl)(nil)

func NewAuthService(passwordHasher password_hasher.PasswordHasher) *AuthServiceImpl {
	return &AuthServiceImpl{passwordHasher}
}

func (s *AuthServiceImpl) VerifyPassword(password string, hash string) error {
	return s.passwordHasher.Verify(password, hash)
}