package expense_infra

import (
	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expense_infra_postgres "github.com/fernan-x/expense-tracker/internal/expense/infra/postgres"
	"github.com/go-pg/pg"
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
	_, err := r.db.Model(expense_infra_postgres.ExpenseModel{}).Insert(expense_infra_postgres.FromDomain(e))
	return err
}

func (r *PostgresExpenseRepository) GetAll() ([]expense_domain.Expense, error) {
	var expenses []expense_infra_postgres.ExpenseModel
	err := r.db.Model(&expenses).Select()
	if err != nil {
		return nil, err
	}

	var res []expense_domain.Expense
	for _, e := range expenses {
		res = append(res, expense_infra_postgres.ToDomain(e))
	}

	return res, nil
}

func (r *PostgresExpenseRepository) Delete(id string) error {
	_, err := r.db.Model(expense_infra_postgres.ExpenseModel{}).Where("id = ?", id).Delete()
	return err
}