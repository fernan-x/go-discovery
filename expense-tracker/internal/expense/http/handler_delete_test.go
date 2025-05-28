package expensehttp_test

import (
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

type mockDeleteExpenseSuccessUseCase struct {}
func (u *mockDeleteExpenseSuccessUseCase) Execute(id string) error {
	return nil
}
type mockDeleteExpenseFailureUseCase struct {}
func (u *mockDeleteExpenseFailureUseCase) Execute(id string) error {
	return fmt.Errorf("not found")
}

func setupDeleteExpenseHandler(uc expenseusecase.DeleteExpenseUseCaseInterface) (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := &expensehttp.ExpenseHandler{
		DeleteExpenseUC: uc,
	}
	router.DELETE("/expenses/:id", handler.DeleteExpense)

	w := httptest.NewRecorder()
	return router, w
}

func TestDeleteExpenseHandler_Success(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/expenses/1", nil)
	req.Header.Set("Content-Type", "application/json")

	router, w := setupDeleteExpenseHandler(&mockDeleteExpenseSuccessUseCase{})
	router.ServeHTTP(w, req)

	var res httpresponse.BaseResponse[expensehttp.DeleteExpenseResponseData]
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", res.Status)
	assert.Equal(t, "1", res.Data.ID)
}

func TestDeleteExpense_Failure_EmptyParameter(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/expenses/1", nil)
	req.Header.Set("Content-Type", "application/json")
	router, w := setupDeleteExpenseHandler(&mockDeleteExpenseFailureUseCase{})
	router.ServeHTTP(w, req)

	// Parse response
	var res httpresponse.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error", res.Status)
	assert.Contains(t, res.Error, "not found")
}