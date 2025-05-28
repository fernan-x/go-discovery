package authusecase

import (
	"fmt"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
)

type RefreshTokenResponse struct {
	AccessToken string
	RefreshToken string
}

type RefreshTokenUseCaseInterface interface {
	Execute(refreshToken string) (RefreshTokenResponse, error)
}

var _ RefreshTokenUseCaseInterface = &RefreshTokenUseCase{}

type RefreshTokenUseCase struct {
	authService authdomain.AuthService
}

func NewRefreshTokenUseCase(authService authdomain.AuthService) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{authService}
}

func (u *RefreshTokenUseCase) Execute(refreshToken string) (RefreshTokenResponse, error) {
	res := RefreshTokenResponse{}
	claims, err := u.authService.ParseToken(refreshToken)
	if err != nil {
		return res, err
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return res, fmt.Errorf("could not parse user id")
	}

	accessToken, err := u.authService.GenerateAccessToken(userId)
	if err != nil {
		return res, fmt.Errorf("could not generate access token: %s", err.Error())
	}

	refreshToken, err = u.authService.GenerateRefreshToken(userId)
	if err != nil {
		return res, fmt.Errorf("could not generate refresh token: %s", err.Error())
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return res, nil
}