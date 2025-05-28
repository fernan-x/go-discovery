package expenseinfra

import (
	"errors"
	"fmt"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expenseinfra_postgres "github.com/fernan-x/expense-tracker/internal/expense/infra/postgres"
	"github.com/go-pg/pg/v10"
)

type PostgresExpenseRepository struct {
	db *pg.DB
}

var _ expensedomain.ExpenseRepository = (*PostgresExpenseRepository)(nil)

func NewPostgresExpenseRepository(db *pg.DB) *PostgresExpenseRepository {
	return &PostgresExpenseRepository{
		db: db,
	}
}

func (r *PostgresExpenseRepository) Create(e expensedomain.Expense) error {
	expense := expenseinfra_postgres.FromDomain(e)
	_, err := r.db.Model(&expense).Insert()
	return err
}

func (r *PostgresExpenseRepository) GetAll() ([]expensedomain.Expense, error) {
	var expenses []expenseinfra_postgres.ExpenseModel
	err := r.db.Model(&expenses).Select()
	if err != nil {
		return nil, err
	}

	var res []expensedomain.Expense = make([]expensedomain.Expense, 0)
	for _, e := range expenses {
		res = append(res, expenseinfra_postgres.ToDomain(e))
	}

	return res, nil
}

func (r *PostgresExpenseRepository) Delete(id string) error {
	res, err := r.db.Model(&expenseinfra_postgres.ExpenseModel{}).Where("id = ?", id).Delete()
	if res.RowsAffected() == 0 {
		return fmt.Errorf("No expense found with id %s", id)
	}
	return err
}

func (r *PostgresExpenseRepository) Update(id string, fields expensedomain.ExpenseUpdateFields) error {
	return errors.New("Not implemented")
}

func (r *PostgresExpenseRepository) GetByID(id string) (*expensedomain.Expense, error) {
	return nil, errors.New("Not implemented")
}