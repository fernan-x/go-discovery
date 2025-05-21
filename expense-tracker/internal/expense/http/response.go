package expense_http

type AddExpenseResponseData struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}