package expenseusecase_test

import (
	"testing"
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expenseinfra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func setupUC() (expenseusecase.UpdateExpenseUseCaseInterface, *expenseinfra.InMemoryExpenseRepository) {
	repo := expenseinfra.NewInMemoryExpenseRepository()
	return expenseusecase.NewUpdateExpenseUseCase(repo), repo
}

func TestUpdateExpense_Failure_TitleRequired(t *testing.T) {
	uc, _ := setupUC()

	emptyTitle := ""
	err := uc.Execute("1", expensedomain.ExpenseUpdateFields{
		Title: &emptyTitle,
	})
	assert.Error(t, err)
	assert.Equal(t, "title cannot be empty", err.Error())
}

func TestUpdateExpense_Failure_AmountLessThanZero(t *testing.T) {
	uc, _ := setupUC()

	negativeAmount := -10.00
	err := uc.Execute("1", expensedomain.ExpenseUpdateFields{
		Amount: &negativeAmount,
	})
	assert.Error(t, err)
	assert.Equal(t, "amount cannot be less than or equal to 0", err.Error())
}

func TestUpdateExpense_Failure_AmountEqualZero(t *testing.T) {
	uc, _ := setupUC()

	negativeAmount := 0.00
	err := uc.Execute("1", expensedomain.ExpenseUpdateFields{
		Amount: &negativeAmount,
	})
	assert.Error(t, err)
	assert.Equal(t, "amount cannot be less than or equal to 0", err.Error())
}

func TestUpdateExpense_Success_TwoFields(t *testing.T) {
	uc, repo := setupUC()

	repo.Create(expensedomain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})

	title := "Lunch edited"
	amount := 22.00
	err := uc.Execute("1", expensedomain.ExpenseUpdateFields{
		Title:   &title,
		Amount:  &amount,
	})
	assert.NoError(t, err)

	expenses, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(expenses))
	assert.Equal(t, "Lunch edited", expenses[0].Title)
	assert.Equal(t, 22.00, expenses[0].Amount)
}

func TestUpdateExpense_Success_OneField(t *testing.T) {
	uc, repo := setupUC()

	repo.Create(expensedomain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})

	title := "Lunch edited"
	err := uc.Execute("1", expensedomain.ExpenseUpdateFields{
		Title:   &title,
	})
	assert.NoError(t, err)

	expenses, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(expenses))
	assert.Equal(t, "Lunch edited", expenses[0].Title)
	assert.Equal(t, 12.90, expenses[0].Amount)
}