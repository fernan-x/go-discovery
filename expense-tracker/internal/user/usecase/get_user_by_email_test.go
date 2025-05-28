package userusecase_test

import (
	"testing"

	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	userinfra "github.com/fernan-x/expense-tracker/internal/user/infra"
	userusecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

var repo = userinfra.NewInMemoryUserRepository()

func TestGetUserByEmailUsecase_Failing(t *testing.T) {
	uc := userusecase.NewGetUserByEmailUseCase(repo)

	_, err := uc.Execute("jean.dupont@test.com")
	assert.Error(t, err)
	assert.Equal(t, "user with email jean.dupont@test.com not found", err.Error())
}

func TestGetUserByEmailUsecase_Success(t *testing.T) {
	ucCreate := userusecase.NewCreateUserUseCase(repo, &passwordhasher.BcryptPasswordHasher{})
	err := ucCreate.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	assert.NoError(t, err)

	uc := userusecase.NewGetUserByEmailUseCase(repo)

	user, err := uc.Execute("jean.dupont@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "jean.dupont@test.com", user.Email)
	assert.Equal(t, "Jean", user.FirstName)
	assert.Equal(t, "Dupont", user.LastName)
}

func TestGetUserByEmailUsecase_Failure_UserNotFound(t *testing.T) {
	ucCreate := userusecase.NewCreateUserUseCase(repo, &passwordhasher.BcryptPasswordHasher{})
	err := ucCreate.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	assert.NoError(t, err)

	uc := userusecase.NewGetUserByEmailUseCase(repo)

	_, err = uc.Execute("jean.dupont2@test.com")
	assert.Error(t, err)
	assert.Equal(t, "user with email jean.dupont2@test.com not found", err.Error())
}