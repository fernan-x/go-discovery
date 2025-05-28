package expensehttp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	expensehttp "github.com/fernan-x/expense-tracker/internal/expense/http"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/httpresponse"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockAddExpenseSuccessUseCase struct {}
func (u *mockAddExpenseSuccessUseCase) Execute(title string, amount float64) error {
	return nil
}
type mockAddExpenseFailureUseCase struct {}
func (u *mockAddExpenseFailureUseCase) Execute(title string, amount float64) error {
	return fmt.Errorf("something went wrong")
}

func setupAddExpenseHandler(uc expenseusecase.AddExpenseUseCaseInterface) (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := &expensehttp.ExpenseHandler{
		AddExpenseUC: uc,
	}
	router.POST("/expenses", handler.AddExpense)

	w := httptest.NewRecorder()
	return router, w
}

func TestAddExpenseHandler_Success(t *testing.T) {
	body := []byte(`{"title": "Lunch", "amount": 12.90}`)
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup
	router, w := setupAddExpenseHandler(&mockAddExpenseSuccessUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.BaseResponse[expensehttp.AddExpenseResponseData]
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, "success", res.Status)
	assert.Equal(t, "Lunch", res.Data.Title)
	assert.Equal(t, 12.9, res.Data.Amount)
}

func TestAddExpenseHandler_Failure_TitleRequired(t *testing.T) {
	body := []byte(`{"amount": 12.90}`)
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router, w := setupAddExpenseHandler(&mockAddExpenseSuccessUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", res.Status)
	assert.Contains(t, res.Error, "Field validation for 'Title'")

}

func TestAddExpenseHandler_Failure_AmountLessThanZero(t *testing.T) {
	body := []byte(`{"title": "Lunch", "amount": -12.90}`)
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup
	router, w := setupAddExpenseHandler(&mockAddExpenseSuccessUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", res.Status)
	assert.Contains(t, res.Error, "Field validation for 'Amount'")
}

func TestAddExpenseHandler_Failure(t *testing.T){
	body := []byte(`{"title": "Lunch", "amount": 12.90}`)
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup
	router, w := setupAddExpenseHandler(&mockAddExpenseFailureUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", res.Status)
	assert.Contains(t, res.Error, "something went wrong")
}