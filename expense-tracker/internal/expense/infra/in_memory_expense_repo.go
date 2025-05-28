package expenseinfra

import (
	"fmt"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type InMemoryExpenseRepository struct {
	expenses []expensedomain.Expense
}

var _ expensedomain.ExpenseRepository = (*InMemoryExpenseRepository)(nil)

func NewInMemoryExpenseRepository() *InMemoryExpenseRepository {
	return &InMemoryExpenseRepository{
		expenses: []expensedomain.Expense{},
	}
}

func (r *InMemoryExpenseRepository) Create(e expensedomain.Expense) error {
	r.expenses = append(r.expenses, e)
	return nil
}

func (r *InMemoryExpenseRepository) GetAll() ([]expensedomain.Expense, error) {
	return r.expenses, nil
}

func (r *InMemoryExpenseRepository) Delete(id string) error {
	for i, e := range r.expenses {
		if e.ID == id {
			// Delete the element at index i
			r.expenses = append(r.expenses[:i], r.expenses[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("element with id %s not found", id)
}

func applyUpdate(e *expensedomain.Expense, fields expensedomain.ExpenseUpdateFields) {
	if fields.Title != nil {
		e.Title = *fields.Title
	}
	if fields.Amount != nil {
		e.Amount = *fields.Amount
	}
}

func (r *InMemoryExpenseRepository) Update(id string, fields expensedomain.ExpenseUpdateFields) error {
	for i, e := range r.expenses {
		if e.ID == id {
			applyUpdate(&e, fields)
			r.expenses[i] = e
			return nil
		}
	}

	return fmt.Errorf("element with id %s not found", id)
}

func (r *InMemoryExpenseRepository) GetByID(id string) (*expensedomain.Expense, error) {
	for _, e := range r.expenses {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, fmt.Errorf("element with id %s not found", id)
}