package transaction

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID                int
	Uuid              string
	Name              string
	Email             string
	Phone             string
	Address           string
	TransactionStatus string
	TransactionTotal  float64
	Shipping          float64
	Status            string
	PaymentURL        string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime
	Details           []TransactionDetail
}

type TransactionDetail struct {
	ID            int
	TransactionID int
	ProductID     int
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
