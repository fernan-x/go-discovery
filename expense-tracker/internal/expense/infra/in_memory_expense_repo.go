package expense_infra

import (
	"errors"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

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

func (r *InMemoryExpenseRepository) Delete(id string) error {
	for i, e := range r.expenses {
		if e.ID == id {
			// Delete the element at index i
			r.expenses = append(r.expenses[:i], r.expenses[i+1:]...)
			return nil
		}
	}

	return errors.New("Element with id " + id + " not found")
}

func applyUpdate(e *expense_domain.Expense, fields expense_domain.ExpenseUpdateFields) {
	if fields.Title != nil {
		e.Title = *fields.Title
	}
	if fields.Amount != nil {
		e.Amount = *fields.Amount
	}
}

func (r *InMemoryExpenseRepository) Update(id string, fields expense_domain.ExpenseUpdateFields) error {
	for i, e := range r.expenses {
		if e.ID == id {
			applyUpdate(&e, fields)
			r.expenses[i] = e
			return nil
		}
	}

	return errors.New("Element with id " + id + " not found")
}