package expense_http

import expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"

type AddExpenseResponseData struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type GetAllExpenseResponseData = []expense_domain.Expense

type DeleteExpenseResponseData struct {
	ID string `json:"id"`
}