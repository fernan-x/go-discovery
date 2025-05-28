package expenseusecase

import expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type GetExpenseByIdUseCaseInterface interface {
	Execute(id string) (*expensedomain.Expense, error)
}

var _ GetExpenseByIdUseCaseInterface = &GetExpenseByIdUseCase{}

type GetExpenseByIdUseCase struct {
	repo expensedomain.ExpenseRepository
}

func NewGetExpenseByIdUseCase(repo expensedomain.ExpenseRepository) *GetExpenseByIdUseCase {
	return &GetExpenseByIdUseCase{repo}
}

func (u *GetExpenseByIdUseCase) Execute(id string) (*expensedomain.Expense, error) {
	return u.repo.GetByID(id)
}