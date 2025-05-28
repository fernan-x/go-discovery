package auth_usecase

import (
	"errors"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type LoginUseCaseInterface interface {
	Execute(email string, password string) (LoginResponse, error)
}

var _ LoginUseCaseInterface = &LoginUseCase{}

type LoginUseCase struct {
	userRepo user_domain.UserRepository
	authService auth_domain.AuthService
}

type LoginResponse struct {
	AccessToken string
	RefreshToken string
}

func NewLoginUseCase(userRepo user_domain.UserRepository, authService auth_domain.AuthService) *LoginUseCase {
	return &LoginUseCase{userRepo, authService}
}

func (u *LoginUseCase) Execute(email string, password string) (LoginResponse, error) {
	user, err := u.userRepo.GetByEmail(email)
	res := LoginResponse{}
	if err != nil {
		return res, errors.New("Invalid credentials")
	}

	err = u.authService.VerifyPassword(password, user.Password)
	if err != nil {
		return res, errors.New("Invalid credentials")
	}

	accessToken, err := u.authService.GenerateAccessToken(user.ID)
	if err != nil {
		return res, errors.New("Could not generate access token: " + err.Error())
	}

	refreshToken, err := u.authService.GenerateRefreshToken(user.ID)
	if err != nil {
		return res, errors.New("Could not generate refresh token: " + err.Error())
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return res, nil
}