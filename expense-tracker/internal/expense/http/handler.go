package expense_http

import (
	"log"
	"net/http"

	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/httpresponse"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	GetAllExpenseUC expense_usecase.GetAllExpenseUseCaseInterface
	AddExpenseUC    expense_usecase.AddExpenseUseCaseInterface
	DeleteExpenseUC expense_usecase.DeleteExpenseUseCaseInterface
}

func (u *ExpenseHandler) DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, httpresponse.NewErrorResponse("Id parameter is required"))
		return
	}

	err := u.DeleteExpenseUC.Execute(id)
	if err != nil {
		log.Printf("Failed to delete expense: %v", err)
		c.JSON(http.StatusNotFound, httpresponse.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, httpresponse.NewSuccessResponse(DeleteExpenseResponseData{
		ID: id,
	}))
}

func (u *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	expenses, err := u.GetAllExpenseUC.Execute()
	if err != nil {
		log.Printf("Failed to get all expenses: %v", err)
		c.JSON(http.StatusInternalServerError, httpresponse.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, httpresponse.NewSuccessResponse(expenses))
}

func (u *ExpenseHandler) AddExpense(c *gin.Context) {
	var request AddExpenseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error": err.Error(),
		})
		return
	}

	if err := u.AddExpenseUC.Execute(request.Title, request.Amount); err != nil {
		log.Printf("Failed to add expense: %v", err)
		c.JSON(http.StatusInternalServerError, httpresponse.NewErrorResponse(err.Error()))
		return
	}

	response := httpresponse.NewSuccessResponse(AddExpenseResponseData{
		Title:  request.Title,
		Amount: request.Amount,
	})
	c.JSON(http.StatusOK, response)
}
