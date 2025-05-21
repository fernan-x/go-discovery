package expense_http

import (
	"fmt"
	"net/http"

	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/httpresponse"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	GetAllExpenseUC expense_usecase.GetAllExpenseUseCaseInterface
	AddExpenseUC    expense_usecase.AddExpenseUseCaseInterface
}

func (u *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	expenses, err := u.GetAllExpenseUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.NewErrorResponse(err.Error()))
		return
	}
	fmt.Printf("Expenses: %v\n", expenses)
	c.JSON(http.StatusOK, httpresponse.NewSuccessResponse(expenses))
}

func (u *ExpenseHandler) AddExpense(c *gin.Context) {
	var request AddExpenseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	if err := u.AddExpenseUC.Execute(request.Title, request.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.NewErrorResponse(err.Error()))
		return
	}

	response := httpresponse.NewSuccessResponse(AddExpenseResponseData{
		Title:  request.Title,
		Amount: request.Amount,
	})
	c.JSON(http.StatusOK, response)
}
