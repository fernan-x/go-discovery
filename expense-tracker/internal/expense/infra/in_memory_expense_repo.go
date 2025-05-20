package expense_infra

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type InMemoryExpenseRepository struct {
	expenses []expense_domain.Expense
}

var _ expense_domain.ExpenseRepository = (*InMemoryExpenseRepository)(nil)

func NewInMemoryExpenseRepository() *InMemoryExpenseRepository {
	return &InMemoryExpenseRepository{
		expenses: []expense_domain.Expense{},
	}
}

func (r *InMemoryExpenseRepository) Create(e expense_domain.Expense) error {
	r.expenses = append(r.expenses, e)
	return nil
}

func (r *InMemoryExpenseRepository) GetAll() ([]expense_domain.Expense, error) {
	return r.expenses, nil
}
