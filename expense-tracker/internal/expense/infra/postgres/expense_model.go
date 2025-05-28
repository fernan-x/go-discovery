package expenseinfra_postgres

import "time"

type ExpenseModel struct {
	tableName struct{} 	 `pg:"expenses"`
	ID        string     `pg:"id,pk,notnull"`
	Title  	  string     `pg:"title,notnull"`
	Amount    float64    `pg:"amount,notnull"`
	CreatedAt time.Time  `pg:"created_at,notnull"`
}