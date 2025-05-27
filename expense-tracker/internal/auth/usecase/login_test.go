package auth_usecase_test

import (
	"testing"

	auth_domain "github.com/fernan-x/expense-tracker/internal/auth/domain"
	auth_usecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	password_hasher "github.com/fernan-x/expense-tracker/internal/password-hasher"
	user_infra "github.com/fernan-x/expense-tracker/internal/user/infra"
	user_usecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

var userRepo = user_infra.NewInMemoryUserRepository()
var passwordHasher = &password_hasher.BcryptPasswordHasher{}
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

	var s auth_domain.AuthService
	uc := auth_usecase.NewLoginUseCase(userRepo, &s)

	_, err := uc.Execute("jean.dupont2@test.com", "123456")
	assert.Error(t, err)
	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestLogin_Success(t *testing.T) {
	initTestData()

	var s auth_domain.AuthService
	uc := auth_usecase.NewLoginUseCase(userRepo, &s)

	token, err := uc.Execute("jean.dupont@test.com", "123456")
	assert.NoError(t, err)
	assert.Equal(t, "xxxx", token)
}

// func TestLogin_Failure_PasswordMissMatch(t *testing.T) {
// 	initTestData()

// 	var s auth_domain.AuthService
// 	uc := auth_usecase.NewLoginUseCase(userRepo, &s)

// 	_, err := uc.Execute("jean.dupont@test.com", "12345")
// 	assert.Error(t, err)
// 	assert.Equal(t, "Invalid credentials", err.Error())
// }