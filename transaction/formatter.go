package transaction

import "time"

type TransactionFormatter struct {
	ID                int       `json:"id"`
	Uuid              string    `json:"uuid"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	TransactionStatus string    `json:"transaction_status"`
	TransactionTotal  float64   `json:"transaction_total"`
	Status            string    `json:"status"`
	PaymentURL        string    `json:"payment_url"`
	CreatedAt         time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	return TransactionFormatter{
		ID:                transaction.ID,
		Uuid:              transaction.Uuid,
		Name:              transaction.Name,
		Email:             transaction.Email,
		TransactionStatus: transaction.Status,
		PaymentURL:        transaction.PaymentURL,
		CreatedAt:         transaction.CreatedAt,
	}
}
