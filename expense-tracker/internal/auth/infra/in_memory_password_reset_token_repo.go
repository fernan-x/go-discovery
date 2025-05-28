package auth_infra

import (
	"errors"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
)


var _ auth_domain.PasswordResetTokenRepository = (*InMemoryPasswordResetTokenRepository)(nil)

type InMemoryPasswordResetTokenRepository struct {
	tokens map[string]auth_domain.PasswordResetToken
}

func NewInMemoryPasswordResetTokenRepository() *InMemoryPasswordResetTokenRepository {
	return &InMemoryPasswordResetTokenRepository{tokens: make(map[string]auth_domain.PasswordResetToken)}
}

func (r *InMemoryPasswordResetTokenRepository) Save(token auth_domain.PasswordResetToken) error {
	r.tokens[token.Token] = token
	return nil
}

func (r *InMemoryPasswordResetTokenRepository) GetByToken(token string) (*auth_domain.PasswordResetToken, error) {
	t, ok := r.tokens[token]
	if !ok {
		return nil, errors.New("Token not found")
	}

	return &t, nil
}