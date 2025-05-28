package expenseusecase

import (
	"errors"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type UpdateExpenseUseCaseInterface interface {
	Execute(id string, input expensedomain.ExpenseUpdateFields) error
}

var _ UpdateExpenseUseCaseInterface = &UpdateExpenseUseCase{}

type UpdateExpenseUseCase struct {
	repo expensedomain.ExpenseRepository
}

func NewUpdateExpenseUseCase(repo expensedomain.ExpenseRepository) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{repo}
}

func (u *UpdateExpenseUseCase) Execute(id string, input expensedomain.ExpenseUpdateFields) error {
	if input.Title != nil && *input.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if input.Amount != nil && *input.Amount <= 0 {
		return errors.New("Amount cannot be less than or equal to 0")
	}

	return u.repo.Update(id, input)
}