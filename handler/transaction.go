package handler

import (
	"fmt"
	"net/http"
	"skripsi-transaction-service/helper"
	"skripsi-transaction-service/transaction"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransactions(c *fiber.Ctx) error {
	var input transaction.CreateTransactionInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse("Failed to create transacation", http.StatusBadRequest, "error", err.Error())
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transacation", http.StatusInternalServerError, "error", err.Error())
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to create transacation", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	return c.JSON(response)
}

func (h *transactionHandler) GetTransaction(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetTransactionByID(id)
	if err != nil {
		response := helper.APIResponse("Failed to get transacation", http.StatusInternalServerError, "error", err.Error())
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	fmt.Println(data)
	response := helper.APIResponse("Success to get transacation", http.StatusOK, "success", transaction.FormatTransaction(data))
	return c.JSON(response)
}

func (h *transactionHandler) GetNotification(c *fiber.Ctx) error {
	var input transaction.TransactionNotificationInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse("Failed to update transacation", http.StatusBadRequest, "error", err.Error())
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to update transacation", http.StatusBadRequest, "error", err.Error())
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	return c.JSON(input)
}
