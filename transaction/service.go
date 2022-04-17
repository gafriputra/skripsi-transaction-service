package transaction

import (
	"errors"
	"skripsi-transaction-service/payment"
	"strconv"
)

type service struct {
	repository     Repository
	paymentService payment.Service
}

type Service interface {
	GetTransactionByID(ID int) (Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, paymentService payment.Service) *service {
	return &service{repository, paymentService}
}

func (s *service) GetTransactionByID(ID int) (Transaction, error) {
	return s.repository.GetByID(ID)
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		Shipping: input.Shipping,
	}

	for _, detail := range input.Details {
		transaction.TransactionTotal += (float64(detail.Quantity) * detail.Price)
	}

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	details := []TransactionDetail{}
	for _, detail := range input.Details {
		details = append(details, TransactionDetail{
			TransactionID: newTransaction.ID,
			ProductID:     detail.ProductID,
			Quantity:      detail.Quantity,
		})
	}

	_, err = s.repository.SaveDetails(details)
	if err != nil {
		return Transaction{}, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.TransactionTotal,
	}
	paymentCustomer := payment.TransactionCustomer{
		Name:  newTransaction.Name,
		Email: newTransaction.Email,
	}
	paymentURL, err := s.paymentService.GetPaymentResponse(paymentTransaction, paymentCustomer)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.Update(newTransaction)
	return newTransaction, err
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	} else if transaction.ID == 0 {
		return errors.New("Transaction Not Found")
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancel"
	}

	_, err = s.repository.Update(transaction)
	if err != nil {
		return err
	}

	return nil

}
