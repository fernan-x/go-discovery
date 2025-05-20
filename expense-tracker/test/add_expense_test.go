package test

import (
	"testing"

	"github.com/fernan-x/expense-tracker/internal/infra"
	"github.com/fernan-x/expense-tracker/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestAddExpense(t *testing.T) {
	repo := infra.NewInMemoryExpenseRepository()
	usecase := usecase.NewAddExpenseUseCase(repo)

	// Add first expense
	err := usecase.Execute("Lunch", 12.90)
	assert.NoError(t, err)

	expenses, err := repo.GetAll()
	assert.Equal(t, 1, len(expenses))
	assert.Equal(t, "Lunch", expenses[0].Title)
	assert.Equal(t, 12.90, expenses[0].Amount)

	// Add a second expense
	err = usecase.Execute("Dinner", 20.00)
	assert.NoError(t, err)

	expenses, err = repo.GetAll()
	assert.Equal(t, 2, len(expenses))
	assert.Equal(t, "Lunch", expenses[0].Title)
	assert.Equal(t, 12.90, expenses[0].Amount)
	assert.Equal(t, "Dinner", expenses[1].Title)
	assert.Equal(t, 20.00, expenses[1].Amount)
}

func TestAddExpenseWithError(t *testing.T) {
	repo := infra.NewFailingExpenseRepositoryTest()
	uc := usecase.NewAddExpenseUseCase(repo)

	err := uc.Execute("Lunch", 12.90)
	assert.Error(t, err)

	expenses, err := repo.GetAll()
	assert.Equal(t, 0, len(expenses))
	assert.Error(t, err)
}