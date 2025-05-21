package expense_usecase

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type DeleteExpenseUseCaseInterface interface {
	Execute(id string) error
}

var _ DeleteExpenseUseCaseInterface = &DeleteExpenseUseCase{}

type DeleteExpenseUseCase struct {
	repo expense_domain.ExpenseRepository
}

func NewDeleteExpenseUseCase(repo expense_domain.ExpenseRepository) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{repo}
}

func (u *DeleteExpenseUseCase) Execute(id string) error {
	return u.repo.Delete(id)
}