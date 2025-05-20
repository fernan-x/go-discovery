package expense_usecase

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type GetAllExpenseUseCase struct {
	repo expense_domain.ExpenseRepository
}

func NewGetAllExpenseUseCase(repo expense_domain.ExpenseRepository) *GetAllExpenseUseCase {
	return &GetAllExpenseUseCase{repo}
}

func (u *GetAllExpenseUseCase) Execute() ([]expense_domain.Expense, error) {
	return u.repo.GetAll()
}