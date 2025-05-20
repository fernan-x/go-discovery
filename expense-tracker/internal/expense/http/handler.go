package expense_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllExpenses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}