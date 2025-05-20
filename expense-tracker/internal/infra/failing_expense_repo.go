package infra

import (
	"errors"

	"github.com/fernan-x/expense-tracker/internal/domain"
)

type FailingExpenseRepositoryTest struct {}

var _ domain.ExpenseRepository = (*FailingExpenseRepositoryTest)(nil)

func NewFailingExpenseRepositoryTest() *FailingExpenseRepositoryTest {
	return &FailingExpenseRepositoryTest{}
}

func (r *FailingExpenseRepositoryTest) Create(e domain.Expense) error {
	return errors.New("Error creating expense")
}

func (r *FailingExpenseRepositoryTest) GetAll() ([]domain.Expense, error) {
	return nil, errors.New("Error getting all expenses")
}