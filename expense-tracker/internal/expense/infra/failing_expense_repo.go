package expense_infra

import (
	"errors"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
)

type FailingExpenseRepositoryTest struct {}

var _ expense_domain.ExpenseRepository = (*FailingExpenseRepositoryTest)(nil)

func NewFailingExpenseRepositoryTest() *FailingExpenseRepositoryTest {
	return &FailingExpenseRepositoryTest{}
}

func (r *FailingExpenseRepositoryTest) Create(e expense_domain.Expense) error {
	return errors.New("Error creating expense")
}

func (r *FailingExpenseRepositoryTest) GetAll() ([]expense_domain.Expense, error) {
	return nil, errors.New("Error getting all expenses")
}