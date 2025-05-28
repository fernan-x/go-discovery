package expensehttp

import expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type AddExpenseResponseData struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type GetAllExpenseResponseData = []expensedomain.Expense

type DeleteExpenseResponseData struct {
	ID string `json:"id"`
}