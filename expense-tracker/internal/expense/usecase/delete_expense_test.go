package expenseusecase_test

import (
	"testing"
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expenseinfra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func TestDeleteExpense_Failure(t *testing.T) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	uc := expenseusecase.NewDeleteExpenseUseCase(repo)

	err := uc.Execute("1")
	assert.Error(t, err)
	assert.Equal(t, "Element with id 1 not found", err.Error())
}

func TestDeleteExpense_Success(t *testing.T) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	uc := expenseusecase.NewDeleteExpenseUseCase(repo)

	err := repo.Create(expensedomain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})
	assert.NoError(t, err)

	err = uc.Execute("1")
	assert.NoError(t, err)

	expenses, err := repo.GetAll()
	assert.Equal(t, 0, len(expenses))
	assert.NoError(t, err)
}