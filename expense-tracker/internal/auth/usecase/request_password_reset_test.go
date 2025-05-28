package authusecase_test

import (
	"testing"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	authinfra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	authusecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	tokenissuer "github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
	userinfra "github.com/fernan-x/expense-tracker/internal/user/infra"
	"github.com/stretchr/testify/assert"
)

func setupPasswordResetTest() (authdomain.PasswordResetTokenRepository, userdomain.UserRepository, tokenissuer.TokenIssuer) {
	tokenRepo := authinfra.NewInMemoryPasswordResetTokenRepository()
	userRepo := userinfra.NewInMemoryUserRepository()
	tokenIssuer := tokenissuer.NewDummyTokenIssuer()
	return tokenRepo, userRepo, tokenIssuer
}

func TestRequestPasswordReset_Failure_EmailNotFound(t *testing.T) {
	tokenRepo, userRepo, tokenIssuer := setupPasswordResetTest()
	uc := authusecase.NewRequestPasswordResetUseCase(tokenRepo, userRepo, tokenIssuer)

	_, err := uc.Execute("jean.dupont@test.com")
	assert.Error(t, err)
	assert.Equal(t, "User with email jean.dupont@test.com not found", err.Error())
}

func TestRequestPasswordReset_Success(t *testing.T) {
	tokenRepo, userRepo, tokenIssuer := setupPasswordResetTest()

	userRepo.Create(&userdomain.User{
		ID: "12345",
		Email: "jean.dupont@test.com",
	})
	uc := authusecase.NewRequestPasswordResetUseCase(tokenRepo, userRepo, tokenIssuer)

	token, err := uc.Execute("jean.dupont@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "random-token", token)
}