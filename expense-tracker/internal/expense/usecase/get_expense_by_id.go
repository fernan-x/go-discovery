package expense_usecase

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type GetExpenseByIdUseCaseInterface interface {
	Execute(id string) (*expense_domain.Expense, error)
}

var _ GetExpenseByIdUseCaseInterface = &GetExpenseByIdUseCase{}

type GetExpenseByIdUseCase struct {
	repo expense_domain.ExpenseRepository
}

func NewGetExpenseByIdUseCase(repo expense_domain.ExpenseRepository) *GetExpenseByIdUseCase {
	return &GetExpenseByIdUseCase{repo}
}

func (u *GetExpenseByIdUseCase) Execute(id string) (*expense_domain.Expense, error) {
	return u.repo.GetByID(id)
}