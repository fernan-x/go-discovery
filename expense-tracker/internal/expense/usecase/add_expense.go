package expenseusecase

import (
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	"github.com/google/uuid"
)

type AddExpenseUseCaseInterface interface {
	Execute(title string, amount float64) error
}
var _ AddExpenseUseCaseInterface = &AddExpenseUseCase{}

type AddExpenseUseCase struct {
	repo expensedomain.ExpenseRepository
}

func NewAddExpenseUseCase(repo expensedomain.ExpenseRepository) *AddExpenseUseCase {
	return &AddExpenseUseCase{repo}
}


// Method with pointer receiver (like method in OOP)
func (u *AddExpenseUseCase) Execute(title string, amount float64) error {
	e := expensedomain.Expense{
		ID: uuid.New().String(),
		Title: title,
		Amount: amount,
		CreatedAt: time.Now(),
	}

	return u.repo.Create(e)
}