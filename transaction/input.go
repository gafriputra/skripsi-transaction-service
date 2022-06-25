package transaction

type CreateTransactionInput struct {
	Name     string                         `json:"name" binding:"required"`
	Email    string                         `json:"email" binding:"required"`
	Phone    string                         `json:"phone" binding:"required"`
	Address  string                         `json:"address" binding:"required"`
	Shipping float64                        `json:"shipping" binding:"required"`
	Details  []CreateTransactionDetailInput `json:"details" binding:"required"`
}

type CreateTransactionDetailInput struct {
	ProductID int     `json:"product_id" binding:"required"`
	Quantity  int     `json:"qty" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
