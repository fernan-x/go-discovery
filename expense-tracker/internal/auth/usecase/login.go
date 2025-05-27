package auth_usecase

import (
	"errors"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type LoginUseCaseInterface interface {
	Execute(email string, password string) (bool, error)
}

var _ LoginUseCaseInterface = &LoginUseCase{}

type LoginUseCase struct {
	userRepo user_domain.UserRepository
	authService auth_domain.AuthService
}

func NewLoginUseCase(userRepo user_domain.UserRepository, authService auth_domain.AuthService) *LoginUseCase {
	return &LoginUseCase{userRepo, authService}
}

func (u *LoginUseCase) Execute(email string, password string) (bool, error) {
	// user, err := u.userRepo.getByEmail(email)
	return false, errors.New("Not implemented")
}