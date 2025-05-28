package userusecase_test

import (
	"testing"

	passwordhasher "github.com/fernan-x/expense-tracker/internal/shared/passwordhasher"
	userinfra "github.com/fernan-x/expense-tracker/internal/user/infra"
	userusecase "github.com/fernan-x/expense-tracker/internal/user/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := userinfra.NewInMemoryUserRepository()
	passwordHasher := &passwordhasher.BcryptPasswordHasher{}
	uc := userusecase.NewCreateUserUseCase(repo, passwordHasher)

	err := uc.Execute("jean.dupont@test.com", "123456", "Jean", "Dupont")
	assert.NoError(t, err)

	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))

	createdUser := users[0]
	assert.Equal(t, "jean.dupont@test.com", createdUser.Email)
	assert.Equal(t, "Jean", createdUser.FirstName)
	assert.Equal(t, "Dupont", createdUser.LastName)

	// Password must be hashed
	assert.NotEqual(t, "123456", createdUser.Password)
	err = passwordHasher.Verify("123456", createdUser.Password)
	assert.NoError(t, err)
}