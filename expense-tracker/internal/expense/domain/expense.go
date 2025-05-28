package expensedomain

import "time"

type Expense struct {
	ID 	        string
	Title       string
	Amount      float64
	CreatedAt   time.Time
}

type ExpenseUpdateFields struct {
	Title *string
	Amount *float64
}

type ExpenseRepository interface {
	Create(e Expense) error
	GetAll() ([]Expense, error)
	Delete(id string) error
	Update(id string, fields ExpenseUpdateFields) error
	GetByID(id string) (*Expense, error)
}