package usecase

import (
	"time"

	"github.com/fernan-x/expense-tracker/internal/domain"
	"github.com/google/uuid"
)

type AddExpenseUseCase struct {
	repo domain.ExpenseRepository
}

func NewAddExpenseUseCase(repo domain.ExpenseRepository) *AddExpenseUseCase {
	return &AddExpenseUseCase{repo}
}

// Method with pointer receiver (like method in OOP)
func (u *AddExpenseUseCase) Execute(title string, amount float64) error {
	e := domain.Expense{
		ID: uuid.New().String(),
		Title: title,
		Amount: amount,
		CreatedAt: time.Now(),
	}

	return u.repo.Create(e)
}