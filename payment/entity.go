package payment

type Transaction struct {
	ID     int
	Amount float64
}

type TransactionCustomer struct {
	Name  string
	Email string
	Phone string
}
