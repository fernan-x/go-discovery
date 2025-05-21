package expense_usecase

import (
	"time"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	"github.com/google/uuid"
)

type AddExpenseUseCaseInterface interface {
	Execute(title string, amount float64) error
}
var _ AddExpenseUseCaseInterface = &AddExpenseUseCase{}

type AddExpenseUseCase struct {
	repo expense_domain.ExpenseRepository
}

func NewAddExpenseUseCase(repo expense_domain.ExpenseRepository) *AddExpenseUseCase {
	return &AddExpenseUseCase{repo}
}


// Method with pointer receiver (like method in OOP)
func (u *AddExpenseUseCase) Execute(title string, amount float64) error {
	e := expense_domain.Expense{
		ID: uuid.New().String(),
		Title: title,
		Amount: amount,
		CreatedAt: time.Now(),
	}

	return u.repo.Create(e)
}