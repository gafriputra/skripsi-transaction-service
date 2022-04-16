package payment

import (
	"errors"
	"skripsi-transaction-service/helper"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct{}

type Service interface {
	GetPaymentResponse(transaction Transaction, customer TransactionCustomer) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentResponse(transaction Transaction, customer TransactionCustomer) (string, error) {
	// 1. Initiate Snap client
	var sn = snap.Client{}
	sn.New(helper.Env("MIDTRANS_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).
	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		// CreditCard: &snap.CreditCardDetails{
		// 	Secure: true,
		// },
		CustomerDetail: &midtrans.CustomerDetails{
			FName: customer.Name,
			LName: "",
			Email: customer.Email,
			Phone: customer.Phone,
		},
	}
	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, err := sn.CreateTransaction(req)
	if err != nil {
		return "", errors.New(err.Message)
	}
	return snapResp.RedirectURL, nil
}
