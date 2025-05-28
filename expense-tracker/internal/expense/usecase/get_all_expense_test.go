package expenseusecase_test

import (
	"testing"
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expenseinfra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetAllExpense(t *testing.T) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	uc := expenseusecase.NewGetAllExpenseUseCase(repo)

	expenses, err := uc.Execute()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(expenses))

	err = repo.Create(expensedomain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	});
	assert.NoError(t, err)

	err = repo.Create(expensedomain.Expense{
		ID: "2",
		Title: "Dinner",
		Amount: 20.00,
		CreatedAt: time.Now(),
	});
	assert.NoError(t, err)

	expenses, err = uc.Execute()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(expenses))
	assert.Equal(t, "Lunch", expenses[0].Title)
	assert.Equal(t, 12.90, expenses[0].Amount)
	assert.Equal(t, "Dinner", expenses[1].Title)
	assert.Equal(t, 20.00, expenses[1].Amount)
}