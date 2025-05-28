package expenseusecase

import expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type GetAllExpenseUseCaseInterface interface {
	Execute() ([]expensedomain.Expense, error)
}

var _ GetAllExpenseUseCaseInterface = &GetAllExpenseUseCase{}

type GetAllExpenseUseCase struct {
	repo expensedomain.ExpenseRepository
}

func NewGetAllExpenseUseCase(repo expensedomain.ExpenseRepository) *GetAllExpenseUseCase {
	return &GetAllExpenseUseCase{repo}
}

func (u *GetAllExpenseUseCase) Execute() ([]expensedomain.Expense, error) {
	return u.repo.GetAll()
}