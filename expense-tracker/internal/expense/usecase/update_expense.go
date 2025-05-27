package expense_usecase

import (
	"errors"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type UpdateExpenseUseCaseInterface interface {
	Execute(id string, input expense_domain.ExpenseUpdateFields) error
}

var _ UpdateExpenseUseCaseInterface = &UpdateExpenseUseCase{}

type UpdateExpenseUseCase struct {
	repo expense_domain.ExpenseRepository
}

func NewUpdateExpenseUseCase(repo expense_domain.ExpenseRepository) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{repo}
}

func (u *UpdateExpenseUseCase) Execute(id string, input expense_domain.ExpenseUpdateFields) error {
	if input.Title != nil && *input.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if input.Amount != nil && *input.Amount <= 0 {
		return errors.New("Amount cannot be less than or equal to 0")
	}

	return u.repo.Update(id, input)
}