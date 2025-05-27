package user_usecase_test

import (
	"testing"

	password_hasher "github.com/fernan-x/expense-tracker/internal/password-hasher"
	user_infra "github.com/fernan-x/expense-tracker/internal/user/infra"
	user_usecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

var repo = user_infra.NewInMemoryUserRepository()

func TestGetUserByEmailUsecase_Failing(t *testing.T) {
	uc := user_usecase.NewGetUserByEmailUseCase(repo)

	_, err := uc.Execute("jean.dupont@test.com")
	assert.Error(t, err)
	assert.Equal(t, "User with email jean.dupont@test.com not found", err.Error())
}

func TestGetUserByEmailUsecase_Success(t *testing.T) {
	ucCreate := user_usecase.NewCreateUserUseCase(repo, &password_hasher.BcryptPasswordHasher{})
	err := ucCreate.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	assert.NoError(t, err)

	uc := user_usecase.NewGetUserByEmailUseCase(repo)

	user, err := uc.Execute("jean.dupont@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "jean.dupont@test.com", user.Email)
	assert.Equal(t, "Jean", user.FirstName)
	assert.Equal(t, "Dupont", user.LastName)
}

func TestGetUserByEmailUsecase_Failure_UserNotFound(t *testing.T) {
	ucCreate := user_usecase.NewCreateUserUseCase(repo, &password_hasher.BcryptPasswordHasher{})
	err := ucCreate.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	assert.NoError(t, err)

	uc := user_usecase.NewGetUserByEmailUseCase(repo)

	_, err = uc.Execute("jean.dupont2@test.com")
	assert.Error(t, err)
	assert.Equal(t, "User with email jean.dupont2@test.com not found", err.Error())
}