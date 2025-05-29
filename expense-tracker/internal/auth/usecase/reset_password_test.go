package authusecase_test

import (
	"testing"

	authusecase "github.com/fernan-x/expense-tracker/internal/auth/usecase"
	"github.com/stretchr/testify/assert"
)

func TestResetPassword_Success(t *testing.T) {
	uc := authusecase.NewResetPasswordUsecase(nil, nil, nil)
	err := uc.Execute("reset-token", "new-password")
	assert.NoError(t, err)
}