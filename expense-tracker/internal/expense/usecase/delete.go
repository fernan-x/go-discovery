package expenseusecase

import expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type DeleteExpenseUseCaseInterface interface {
	Execute(id string) error
}

var _ DeleteExpenseUseCaseInterface = &DeleteExpenseUseCase{}

type DeleteExpenseUseCase struct {
	repo expensedomain.ExpenseRepository
}

func NewDeleteExpenseUseCase(repo expensedomain.ExpenseRepository) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{repo}
}

func (u *DeleteExpenseUseCase) Execute(id string) error {
	return u.repo.Delete(id)
}