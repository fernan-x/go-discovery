package expenseusecase

import (
	"fmt"

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
		return fmt.Errorf("title cannot be empty")
	}

	if input.Amount != nil && *input.Amount <= 0 {
		return fmt.Errorf("amount cannot be less than or equal to 0")
	}

	return u.repo.Update(id, input)
}