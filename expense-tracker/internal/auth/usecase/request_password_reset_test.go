package auth_usecase_test

import (
	"testing"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	auth_infra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	auth_usecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	token_issuer "github.com/fernan-x/expense-tracker/internal/shared/token-issuer"
	user_domain "github.com/fernan-x/expense-tracker/internal/user/domain"
	user_infra "github.com/fernan-x/expense-tracker/internal/user/infra"
	"github.com/stretchr/testify/assert"
)

func setupPasswordResetTest() (auth_domain.PasswordResetTokenRepository, user_domain.UserRepository, token_issuer.TokenIssuer) {
	tokenRepo := auth_infra.NewInMemoryPasswordResetTokenRepository()
	userRepo := user_infra.NewInMemoryUserRepository()
	tokenIssuer := token_issuer.NewDummyTokenIssuer()
	return tokenRepo, userRepo, tokenIssuer
}

func TestRequestPasswordReset_Failure_EmailNotFound(t *testing.T) {
	tokenRepo, userRepo, tokenIssuer := setupPasswordResetTest()
	uc := auth_usecase.NewRequestPasswordResetUseCase(tokenRepo, userRepo, tokenIssuer)

	_, err := uc.Execute("jean.dupont@test.com")
	assert.Error(t, err)
	assert.Equal(t, "User with email jean.dupont@test.com not found", err.Error())
}

func TestRequestPasswordReset_Success(t *testing.T) {
	tokenRepo, userRepo, tokenIssuer := setupPasswordResetTest()

	userRepo.Create(&user_domain.User{
		ID: "12345",
		Email: "jean.dupont@test.com",
	})
	uc := auth_usecase.NewRequestPasswordResetUseCase(tokenRepo, userRepo, tokenIssuer)

	token, err := uc.Execute("jean.dupont@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "random-token", token)
}