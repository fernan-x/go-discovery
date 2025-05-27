package expense_usecase_test

import (
	"testing"
	"time"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expense_infra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseById_Failure(t *testing.T) {
	repo := expense_infra.NewInMemoryExpenseRepository()
	uc := expense_usecase.NewGetExpenseByIdUseCase(repo)

	e, err := uc.Execute("1")
	assert.Equal(t, "Element with id 1 not found", err.Error())
	assert.Error(t, err)
	assert.Nil(t, e)
}

func TestGetExpenseById_Success(t *testing.T) {
	repo := expense_infra.NewInMemoryExpenseRepository()
	uc := expense_usecase.NewGetExpenseByIdUseCase(repo)

	repo.Create(expense_domain.Expense{
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