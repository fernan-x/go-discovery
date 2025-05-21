package main

import (
	"log"
	"net/http"
	"os"

	expense_http "github.com/fernan-x/expense-tracker/internal/expense/http"
	expense_infra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expense_usecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "youruser",
		Password: "yourpass",
		Database: "yourdb",
	})

	// Ping to test
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		os.Exit(1)
	}

	return db
}

func main() {
	db := connectDB()
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// expenseRepo := expense_infra.NewInMemoryExpenseRepository()
	expenseRepo := expense_infra.NewPostgresExpenseRepository(db)
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