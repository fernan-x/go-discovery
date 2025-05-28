package authusecase

import (
	"time"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	tokenissuer "github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type RequestPasswordResetUseCaseInterface interface {
	Execute(email string) (string, error)
}

var _ RequestPasswordResetUseCaseInterface = (*RequestPasswordResetUseCase)(nil)

type RequestPasswordResetUseCase struct {
	tokenRepo authdomain.PasswordResetTokenRepository
	userRepo userdomain.UserRepository
	tokenIssuer tokenissuer.TokenIssuer
}

func NewRequestPasswordResetUseCase(
	tokenRepo authdomain.PasswordResetTokenRepository,
	userRepo userdomain.UserRepository,
	tokenIssuer tokenissuer.TokenIssuer,
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

	model := authdomain.PasswordResetToken{
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