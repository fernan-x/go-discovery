package expenseinfra

import (
	"fmt"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type FailingExpenseRepositoryTest struct {}

var _ expensedomain.ExpenseRepository = (*FailingExpenseRepositoryTest)(nil)

func NewFailingExpenseRepositoryTest() *FailingExpenseRepositoryTest {
	return &FailingExpenseRepositoryTest{}
}

func (r *FailingExpenseRepositoryTest) Create(e expensedomain.Expense) error {
	return fmt.Errorf("error creating expense")
}

func (r *FailingExpenseRepositoryTest) GetAll() ([]expensedomain.Expense, error) {
	return nil, fmt.Errorf("error getting all expenses")
}

func (r *FailingExpenseRepositoryTest) Delete(id string) error {
	return fmt.Errorf("error deleting expense")
}

func (r *FailingExpenseRepositoryTest) Update(id string, fields expensedomain.ExpenseUpdateFields) error {
	return fmt.Errorf("error updating expense")
}

func (r *FailingExpenseRepositoryTest) GetByID(id string) (*expensedomain.Expense, error) {
	return nil, fmt.Errorf("error getting expense by id")
}