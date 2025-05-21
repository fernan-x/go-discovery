package main

import (
	"net/http"

	expense_http "github.com/fernan-x/expense-tracker/internal/expense/http"
	expense_infra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	expenseRepo := expense_infra.NewInMemoryExpenseRepository()
	getAllExpenseUC := expense_usecase.NewGetAllExpenseUseCase(expenseRepo)
	addExpenseUC := expense_usecase.NewAddExpenseUseCase(expenseRepo)
	deleteExpenseUC := expense_usecase.NewDeleteExpenseUseCase(expenseRepo)
	handler := expense_http.ExpenseHandler{
		GetAllExpenseUC: getAllExpenseUC,
		AddExpenseUC: addExpenseUC,
		DeleteExpenseUC: deleteExpenseUC,
	}

	router.GET("/expenses", handler.GetAllExpenses)
	router.POST("/expenses", handler.AddExpense)
	router.DELETE("/expenses/:id", handler.DeleteExpense)

	router.Run()
}