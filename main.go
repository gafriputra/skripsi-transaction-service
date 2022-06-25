package main

import (
	"fmt"
	"log"
	"skripsi-transaction-service/handler"
	"skripsi-transaction-service/helper"
	"skripsi-transaction-service/payment"
	"skripsi-transaction-service/transaction"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		helper.Env("DB_HOST"), helper.Env("DB_USER"), helper.Env("DB_PASSWORD"), helper.Env("DB_NAME"), helper.Env("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	paymentService := payment.NewService()

	transactioinRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactioinRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := fiber.New()
	api := app.Group("/api/v1/transaction")
	app.Use(cors.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Get("/:id", transactionHandler.GetTransaction)
	api.Post("/", transactionHandler.CreateTransactions)
	api.Post("/midtrans", transactionHandler.GetNotification)

	app.Listen("0.0.0.0:" + helper.Env("APP_PORT"))
}
