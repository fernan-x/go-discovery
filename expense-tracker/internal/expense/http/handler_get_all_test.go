package expense_http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	expense_domain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expense_http "github.com/fernan-x/expense-tracker/internal/expense/http"
	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/fernan-x/expense-tracker/internal/shared/httpresponse"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockGetAllExpenseSuccessUseCase struct {}
func (u *mockGetAllExpenseSuccessUseCase) Execute() ([]expense_domain.Expense, error) {
	return []expense_domain.Expense{
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
func (u *mockGetAllExpenseFailureUseCase) Execute() ([]expense_domain.Expense, error) {
	return nil, errors.New("something went wrong")
}

func setupGetAllExpenseHandler(uc expense_usecase.GetAllExpenseUseCaseInterface) (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := &expense_http.ExpenseHandler{
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
	var res httpresponse.BaseResponse[expense_http.GetAllExpenseResponseData]
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