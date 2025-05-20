package infra

import "github.com/fernan-x/expense-tracker/internal/domain"

type InMemoryExpenseRepository struct {
	expenses []domain.Expense
}

var _ domain.ExpenseRepository = (*InMemoryExpenseRepository)(nil)

func NewInMemoryExpenseRepository() *InMemoryExpenseRepository {
	return &InMemoryExpenseRepository{
		expenses: []domain.Expense{},
	}
}

func (r *InMemoryExpenseRepository) Create(e domain.Expense) error {
	r.expenses = append(r.expenses, e)
	return nil
}

func (r *InMemoryExpenseRepository) GetAll() ([]domain.Expense, error) {
	return r.expenses, nil
}

