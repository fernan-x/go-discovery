package main

import (
	"log"
	"net/http"
	"os"

	expensedomain "github.com/fernan-x/expense-tracker/internal/expense/domain"
	expensehttp "github.com/fernan-x/expense-tracker/internal/expense/http"
	expenseinfra "github.com/fernan-x/expense-tracker/internal/expense/infra"
	expenseusecase "github.com/fernan-x/expense-tracker/internal/expense/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	// Ping to test
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Connected to DB")

	return db
}

func initRepositories(dummy bool) expensedomain.ExpenseRepository {
	if dummy {
		return expenseinfra.NewInMemoryExpenseRepository()
	}

	db := connectDB()
	InitMigrations()
	return expenseinfra.NewPostgresExpenseRepository(db)
}

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	expenseRepo := initRepositories(true)

	getAllExpenseUC := expenseusecase.NewGetAllExpenseUseCase(expenseRepo)
	addExpenseUC := expenseusecase.NewAddExpenseUseCase(expenseRepo)
	deleteExpenseUC := expenseusecase.NewDeleteExpenseUseCase(expenseRepo)
	handler := expensehttp.ExpenseHandler{
		GetAllExpenseUC: getAllExpenseUC,
		AddExpenseUC: addExpenseUC,
		DeleteExpenseUC: deleteExpenseUC,
	}

	router.GET("/expenses", handler.GetAllExpenses)
	router.POST("/expenses", handler.AddExpense)
	router.DELETE("/expenses/:id", handler.DeleteExpense)

	router.Run()
}