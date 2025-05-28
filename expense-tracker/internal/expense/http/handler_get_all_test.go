package expensehttp_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expensehttp "github.com/fernan-x/expense-tracker/internal/expense/http"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/httpresponse"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockGetAllExpenseSuccessUseCase struct {}
func (u *mockGetAllExpenseSuccessUseCase) Execute() ([]expensedomain.Expense, error) {
	return []expensedomain.Expense{
		{
			ID: "1",
			Title: "Lunch",
			Amount: 12.90,
			CreatedAt: time.Now(),
		},
		{
			ID: "2",
			Title: "Dinner",
			Amount: 20.00,
			CreatedAt: time.Now(),
		},
	}, nil
}
type mockGetAllExpenseFailureUseCase struct {}
func (u *mockGetAllExpenseFailureUseCase) Execute() ([]expensedomain.Expense, error) {
	return nil, fmt.Errorf("something went wrong")
}

func setupGetAllExpenseHandler(uc expenseusecase.GetAllExpenseUseCaseInterface) (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := &expensehttp.ExpenseHandler{
		GetAllExpenseUC: uc,
	}
	router.GET("/expenses", handler.GetAllExpenses)

	w := httptest.NewRecorder()
	return router, w
}

func TestGetAllExpenses_Success(t *testing.T) {
	req, _ := http.NewRequest("GET", "/expenses", nil)
	req.Header.Set("Content-Type", "application/json")

	// Setup
	router, w := setupGetAllExpenseHandler(&mockGetAllExpenseSuccessUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.BaseResponse[expensehttp.GetAllExpenseResponseData]
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", res.Status)
	assert.Equal(t, 2, len(res.Data))
}

func TestGetAllExpenses_Error(t *testing.T) {
	req, _ := http.NewRequest("GET", "/expenses", nil)
	req.Header.Set("Content-Type", "application/json")

	// Setup
	router, w := setupGetAllExpenseHandler(&mockGetAllExpenseFailureUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", res.Status)
	assert.Equal(t, "something went wrong", res.Error)
}