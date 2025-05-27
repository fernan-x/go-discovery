package expense_infra

import (
	"errors"
	"fmt"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expense_infra_postgres "github.com/fernan-x/expense-tracker/internal/expense/infra/postgres"
	"github.com/go-pg/pg/v10"
)

type PostgresExpenseRepository struct {
	db *pg.DB
}

var _ expense_domain.ExpenseRepository = (*PostgresExpenseRepository)(nil)

func NewPostgresExpenseRepository(db *pg.DB) *PostgresExpenseRepository {
	return &PostgresExpenseRepository{
		db: db,
	}
}

func (r *PostgresExpenseRepository) Create(e expense_domain.Expense) error {
	expense := expense_infra_postgres.FromDomain(e)
	_, err := r.db.Model(&expense).Insert()
	return err
}

func (r *PostgresExpenseRepository) GetAll() ([]expense_domain.Expense, error) {
	var expenses []expense_infra_postgres.ExpenseModel
	err := r.db.Model(&expenses).Select()
	if err != nil {
		return nil, err
	}

	var res []expense_domain.Expense = make([]expense_domain.Expense, 0)
	for _, e := range expenses {
		res = append(res, expense_infra_postgres.ToDomain(e))
	}

	return res, nil
}

func (r *PostgresExpenseRepository) Delete(id string) error {
	res, err := r.db.Model(&expense_infra_postgres.ExpenseModel{}).Where("id = ?", id).Delete()
	if res.RowsAffected() == 0 {
		return fmt.Errorf("No expense found with id %s", id)
	}
	return err
}

func (r *PostgresExpenseRepository) Update(id string, fields expense_domain.ExpenseUpdateFields) error {
	return errors.New("Not implemented")
}