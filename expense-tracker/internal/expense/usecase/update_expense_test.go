package expense_usecase_test

import (
	"testing"
	"time"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expense_infra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/stretchr/testify/assert"
)

func setupUC() (expense_usecase.UpdateExpenseUseCaseInterface, *expense_infra.InMemoryExpenseRepository) {
	repo := expense_infra.NewInMemoryExpenseRepository()
	return expense_usecase.NewUpdateExpenseUseCase(repo), repo
}

func TestUpdateExpense_Failure_TitleRequired(t *testing.T) {
	uc, _ := setupUC()

	emptyTitle := ""
	err := uc.Execute("1", expense_domain.ExpenseUpdateFields{
		Title: &emptyTitle,
	})
	assert.Error(t, err)
	assert.Equal(t, "Title cannot be empty", err.Error())
}

func TestUpdateExpense_Failure_AmountLessThanZero(t *testing.T) {
	uc, _ := setupUC()

	negativeAmount := -10.00
	err := uc.Execute("1", expense_domain.ExpenseUpdateFields{
		Amount: &negativeAmount,
	})
	assert.Error(t, err)
	assert.Equal(t, "Amount cannot be less than or equal to 0", err.Error())
}

func TestUpdateExpense_Failure_AmountEqualZero(t *testing.T) {
	uc, _ := setupUC()

	negativeAmount := 0.00
	err := uc.Execute("1", expense_domain.ExpenseUpdateFields{
		Amount: &negativeAmount,
	})
	assert.Error(t, err)
	assert.Equal(t, "Amount cannot be less than or equal to 0", err.Error())
}

func TestUpdateExpense_Success_TwoFields(t *testing.T) {
	uc, repo := setupUC()

	repo.Create(expense_domain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})

	title := "Lunch edited"
	amount := 22.00
	err := uc.Execute("1", expense_domain.ExpenseUpdateFields{
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

	repo.Create(expense_domain.Expense{
		ID: "1",
		Title: "Lunch",
		Amount: 12.90,
		CreatedAt: time.Now(),
	})

	title := "Lunch edited"
	err := uc.Execute("1", expense_domain.ExpenseUpdateFields{
		Title:   &title,
	})
	assert.NoError(t, err)

	expenses, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(expenses))
	assert.Equal(t, "Lunch edited", expenses[0].Title)
	assert.Equal(t, 12.90, expenses[0].Amount)
}