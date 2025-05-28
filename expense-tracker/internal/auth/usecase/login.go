package authusecase

import (
	"fmt"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
)

type LoginUseCaseInterface interface {
	Execute(email string, password string) (LoginResponse, error)
}

var _ LoginUseCaseInterface = &LoginUseCase{}

type LoginUseCase struct {
	userRepo userdomain.UserRepository
	authService authdomain.AuthService
}

type LoginResponse struct {
	AccessToken string
	RefreshToken string
}

func NewLoginUseCase(userRepo userdomain.UserRepository, authService authdomain.AuthService) *LoginUseCase {
	return &LoginUseCase{userRepo, authService}
}

func (u *LoginUseCase) Execute(email string, password string) (LoginResponse, error) {
	user, err := u.userRepo.GetByEmail(email)
	res := LoginResponse{}
	if err != nil {
		return res, fmt.Errorf("invalid credentials")
	}

	err = u.authService.VerifyPassword(password, user.Password)
	if err != nil {
		return res, fmt.Errorf("invalid credentials")
	}

	accessToken, err := u.authService.GenerateAccessToken(user.ID)
	if err != nil {
		return res, fmt.Errorf("could not generate access token: %s", err.Error())
	}

	refreshToken, err := u.authService.GenerateRefreshToken(user.ID)
	if err != nil {
		return res, fmt.Errorf("could not generate refresh token: %s", err.Error())
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return res, nil
}