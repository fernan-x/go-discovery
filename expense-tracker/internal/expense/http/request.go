package expense_http

type AddExpenseRequest struct {
	Title string `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gte=0"`
}