package authinfra

import (
	"fmt"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
)


var _ authdomain.PasswordResetTokenRepository = (*InMemoryPasswordResetTokenRepository)(nil)

type InMemoryPasswordResetTokenRepository struct {
	tokens map[string]authdomain.PasswordResetToken
}

func NewInMemoryPasswordResetTokenRepository() *InMemoryPasswordResetTokenRepository {
	return &InMemoryPasswordResetTokenRepository{tokens: make(map[string]authdomain.PasswordResetToken)}
}

func (r *InMemoryPasswordResetTokenRepository) Save(token authdomain.PasswordResetToken) error {
	r.tokens[token.Token] = token
	return nil
}

func (r *InMemoryPasswordResetTokenRepository) GetByToken(token string) (*authdomain.PasswordResetToken, error) {
	t, ok := r.tokens[token]
	if !ok {
		return nil, fmt.Errorf("token not found")
	}

	return &t, nil
}