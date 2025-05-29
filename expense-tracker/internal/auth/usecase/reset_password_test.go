package authusecase_test

import (
	"testing"
	"time"

	authdomain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	authinfra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	authusecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	"github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
	userdomain "github.com/fernan-x/expense-tracker/internal/user/domain"
	userinfra "github.com/fernan-x/expense-tracker/internal/user/infra"
	"github.com/stretchr/testify/assert"
)

type TestCtx struct {
	pwdHasher *passwordhasher.BcryptPasswordHasher
	tokenIssuer *tokenissuer.DummyTokenIssuer
	authService *authinfra.AuthServiceImpl
	userRepo *userinfra.InMemoryUserRepository
	tokenRepo *authinfra.InMemoryPasswordResetTokenRepository
}

func setupResetPasswordTest() *TestCtx {
	pwdHasher := &passwordhasher.BcryptPasswordHasher{}
	tokenIssuer := tokenissuer.NewDummyTokenIssuer()

	return &TestCtx{
		pwdHasher: pwdHasher,
		tokenIssuer: tokenIssuer,
		authService: authinfra.NewAuthService(pwdHasher, tokenIssuer),
		userRepo: userinfra.NewInMemoryUserRepository(),
		tokenRepo: authinfra.NewInMemoryPasswordResetTokenRepository(),
	}
}

func TestResetPassword_Failure_InvalidToken(t *testing.T) {
	ctx := setupResetPasswordTest()
	uc := authusecase.NewResetPasswordUsecase(ctx.authService, ctx.userRepo, ctx.tokenRepo, ctx.pwdHasher)
	err := uc.Execute("reset-token", "new-password")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid reset token")
}

func TestResetPassword_Failure_Success(t *testing.T) {
	ctx := setupResetPasswordTest()
	ctx.userRepo.Create(&userdomain.User{
		ID: "user-id",
		Email: "user@email.com",
		Password: "user-password",
		FirstName: "user-first-name",
		LastName: "user-last-name",
	})
	ctx.tokenRepo.Save(authdomain.PasswordResetToken{
		Token: "reset-token",
		UserId: "user-id",
		ExpiresAt: time.Now().Add(time.Hour),
	})
	uc := authusecase.NewResetPasswordUsecase(ctx.authService, ctx.userRepo, ctx.tokenRepo, ctx.pwdHasher)
	err := uc.Execute("reset-token", "new-password")
	assert.NoError(t, err)
}