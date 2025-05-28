package auth_usecase

import (
	"time"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	token_issuer "github.com/fernan-x/expense-tracker/internal/shared/token-issuer"
	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type RequestPasswordResetUseCaseInterface interface {
	Execute(email string) (string, error)
}

var _ RequestPasswordResetUseCaseInterface = (*RequestPasswordResetUseCase)(nil)

type RequestPasswordResetUseCase struct {
	tokenRepo auth_domain.PasswordResetTokenRepository
	userRepo user_domain.UserRepository
	tokenIssuer token_issuer.TokenIssuer
}

func NewRequestPasswordResetUseCase(
	tokenRepo auth_domain.PasswordResetTokenRepository,
	userRepo user_domain.UserRepository,
	tokenIssuer token_issuer.TokenIssuer,
) *RequestPasswordResetUseCase {
	return &RequestPasswordResetUseCase{tokenRepo, userRepo, tokenIssuer}
}

func (u *RequestPasswordResetUseCase) Execute(email string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	token, err := u.tokenIssuer.GenerateRandomToken(32)
	if err != nil {
		return "", err
	}

	model := auth_domain.PasswordResetToken{
		Token: token,
		UserId: user.ID,
		ExpiresAt: time.Now().Add(time.Minute * 15), // 15 minutes
	}

	err = u.tokenRepo.Save(model)
	if err != nil {
		return "", err
	}

	return token, nil
}