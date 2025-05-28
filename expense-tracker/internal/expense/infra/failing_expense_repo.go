package expenseinfra

import (
	"errors"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type FailingExpenseRepositoryTest struct {}

var _ expensedomain.ExpenseRepository = (*FailingExpenseRepositoryTest)(nil)

func NewFailingExpenseRepositoryTest() *FailingExpenseRepositoryTest {
	return &FailingExpenseRepositoryTest{}
}

func (r *FailingExpenseRepositoryTest) Create(e expensedomain.Expense) error {
	return errors.New("Error creating expense")
}

func (r *FailingExpenseRepositoryTest) GetAll() ([]expensedomain.Expense, error) {
	return nil, errors.New("Error getting all expenses")
}

func (r *FailingExpenseRepositoryTest) Delete(id string) error {
	return errors.New("Error deleting expense")
}

func (r *FailingExpenseRepositoryTest) Update(id string, fields expensedomain.ExpenseUpdateFields) error {
	return errors.New("Error updating expense")
}

func (r *FailingExpenseRepositoryTest) GetByID(id string) (*expensedomain.Expense, error) {
	return nil, errors.New("Error getting expense by id")
}