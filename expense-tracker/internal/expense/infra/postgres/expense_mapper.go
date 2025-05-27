package expense_infra_postgres

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

func FromDomain(e expense_domain.Expense) ExpenseModel {
	return ExpenseModel{
		ID:        e.ID,
		Title:     e.Title,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
	}
}

func ToDomain(e ExpenseModel) expense_domain.Expense {
	return expense_domain.Expense{
		ID:        e.ID,
		Title:     e.Title,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
	}
}
