package expenseusecase_test

import (
	"testing"
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expenseinfra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseById_Failure(t *testing.T) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	uc := expenseusecase.NewGetExpenseByIdUseCase(repo)

	e, err := uc.Execute("1")
	assert.Equal(t, "element with id 1 not found", err.Error())
	assert.Error(t, err)
	assert.Nil(t, e)
}

func TestGetExpenseById_Success(t *testing.T) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	uc := expenseusecase.NewGetExpenseByIdUseCase(repo)

	repo.Create(expensedomain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})

	e, err := uc.Execute("1")
	assert.NoError(t, err)
	assert.Equal(t, "Lunch", e.Title)
	assert.Equal(t, 12.90, e.Amount)
}