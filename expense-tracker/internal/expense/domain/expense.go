package expense_domain

import "time"

type Expense struct {
	ID 	        string
	Title       string
	Amount      float64
	CreatedAt   time.Time
}

type ExpenseRepository interface {
	Create(e Expense) error
	GetAll() ([]Expense, error)
	Delete(id string) error
}