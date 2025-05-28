package authusecase_test

import (
	"testing"

	authinfra "github.com/fernan-x/expense-tracker/internal/auth/infra"
	authusecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	tokenissuer "github.com/fernan-x/expense-tracker/internal/shared/tokenissuer"
	userinfra "github.com/fernan-x/expense-tracker/internal/user/infra"
	userusecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

var userRepo = userinfra.NewInMemoryUserRepository()
var passwordHasher = &passwordhasher.BcryptPasswordHasher{}
var tokenIssuer = tokenissuer.NewJwtTokenIssuer([]byte("secret"))
var authService = authinfra.NewAuthService(passwordHasher, tokenIssuer)
var isInit = false

// Insert first user in repository only once
func initTestData() {
	if isInit {
		return
	}
	uc := userusecase.NewCreateUserUseCase(userRepo, passwordHasher)
	err := uc.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	if err != nil {
		panic(err)
	}
	isInit = true
}

func TestLogin_Failure_NotFound(t *testing.T) {
	initTestData()

	uc := authusecase.NewLoginUseCase(userRepo, authService)

	_, err := uc.Execute("jean.dupont2@test.com", "123456")
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestLogin_Failure_PasswordMissMatch(t *testing.T) {
	initTestData()

	uc := authusecase.NewLoginUseCase(userRepo, authService)

	_, err := uc.Execute("jean.dupont@test.com", "12345")
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestLogin_Success(t *testing.T) {
	initTestData()

	uc := authusecase.NewLoginUseCase(userRepo, authService)

	token, err := uc.Execute("jean.dupont@test.com", "123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}