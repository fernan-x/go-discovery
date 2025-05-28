package auth_usecase_test

import (
	"testing"

	auth_infra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	auth_usecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	password_hasher "github.com/fernan-x/expense-tracker/internal/shared/password-hasher"
	token_issuer "github.com/fernan-x/expense-tracker/internal/shared/token-issuer"
	user_infra "github.com/fernan-x/expense-tracker/internal/user/infra"
	user_usecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

var userRepo = user_infra.NewInMemoryUserRepository()
var passwordHasher = &password_hasher.BcryptPasswordHasher{}
var tokenIssuer = token_issuer.NewJwtTokenIssuer([]byte("secret"))
var authService = auth_infra.NewAuthService(passwordHasher, tokenIssuer)
var isInit = false

// Insert first user in repository only once
func initTestData() {
	if isInit {
		return
	}
	uc := user_usecase.NewCreateUserUseCase(userRepo, passwordHasher)
	err := uc.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	if err != nil {
		panic(err)
	}
	isInit = true
}

func TestLogin_Failure_NotFound(t *testing.T) {
	initTestData()

	uc := auth_usecase.NewLoginUseCase(userRepo, authService)

	_, err := uc.Execute("jean.dupont2@test.com", "123456")
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestLogin_Failure_PasswordMissMatch(t *testing.T) {
	initTestData()

	uc := auth_usecase.NewLoginUseCase(userRepo, authService)

	_, err := uc.Execute("jean.dupont@test.com", "12345")
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestLogin_Success(t *testing.T) {
	initTestData()

	uc := auth_usecase.NewLoginUseCase(userRepo, authService)

	token, err := uc.Execute("jean.dupont@test.com", "123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}