package main

import (
	"fmt"
	"os"

	"github.com/fernan-x/expense-tracker/internal/infra"
	"github.com/fernan-x/expense-tracker/internal/usecase"
)

func main() {
	// repo := infra.NewInMemoryExpenseRepository()
	repo := infra.NewFailingExpenseRepositoryTest()
	uc := usecase.NewAddExpenseUseCase(repo)

	err := uc.Execute("Lunch", 12.90)
	if err != nil {
		fmt.Printf("Error adding expense: %v\n", err)
		os.Exit(1)
	}

	err = uc.Execute("Dinner", 20.00)
	if err != nil {
		fmt.Printf("Error adding expense: %v\n", err)
		os.Exit(1)
	}

	all, err := repo.GetAll()
	if err != nil {
		fmt.Printf("Error getting expenses: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(all)
}