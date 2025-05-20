package main

import (
	"fmt"

	"github.com/fernan-x/expense-tracker/internal/infra"
	"github.com/fernan-x/expense-tracker/internal/usecase"
)

func main() {
	repo := infra.NewInMemoryExpenseRepository()
	uc := usecase.NewAddExpenseUseCase(repo)

	_ = uc.Execute("Lunch", 12.90)
	_ = uc.Execute("Dinner", 20.00)

	all, _ := repo.GetAll()
	fmt.Println(all)
}