package expense_http

import (
	"net/http"

	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	GetAllExpenseUC *expense_usecase.GetAllExpenseUseCase
	AddExpenseUC    *expense_usecase.AddExpenseUseCase
}

func (u *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	expenses, err := u.GetAllExpenseUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func (u *ExpenseHandler) AddExpense(c *gin.Context) {
	// TODO validation
	title := "Lunch"
	amount := 12.90
	err := u.AddExpenseUC.Execute(title, amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"amount": amount,
	})
}
