package main

import (
	"net/http"

	expense_http "github.com/fernan-x/expense-tracker/internal/expense/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/expenses", expense_http.GetAllExpenses)

	router.Run()
}