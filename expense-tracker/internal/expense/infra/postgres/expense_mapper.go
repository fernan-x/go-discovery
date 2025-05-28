package expenseinfra_postgres

import expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"

func FromDomain(e expensedomain.Expense) ExpenseModel {
	return ExpenseModel{
		ID:        e.ID,
		Title:     e.Title,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
	}
}

func ToDomain(e ExpenseModel) expensedomain.Expense {
	return expensedomain.Expense{
		ID:        e.ID,
		Title:     e.Title,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
	}
}
