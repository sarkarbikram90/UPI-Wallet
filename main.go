package main

import (
	"fmt"
	"time"
)

// User represents a user of the application
type User struct {
	ID           int
	PhoneNumber  string
	Password     string
	Name         string
	Email        string
	AccountMoney float64
}

// Transaction represents a payment transaction
type Transaction struct {
	ID              int
	FromUserID      int
	ToUserID        int
	Amount          float64
	TransactionType string
	CreatedAt       time.Time
}

// PaymentMethod represents different payment methods
type PaymentMethod string

const (
	Card       PaymentMethod = "Card"
	UPI        PaymentMethod = "UPI"
	Netbanking PaymentMethod = "Netbanking"
)

// Application represents the payment application
type Application struct {
	Users        []User
	Transactions []Transaction
}

// RegisterUser allows users to create accounts via their phone number and password
func (app *Application) RegisterUser(phoneNumber, password string) int {
	userID := len(app.Users) + 1
	user := User{
		ID:          userID,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
	app.Users = append(app.Users, user)
	return userID
}

// UpdateUser allows users to update their profile details
func (app *Application) UpdateUser(userID int, name, email, phoneNumber string) {
	for i := range app.Users {
		if app.Users[i].ID == userID {
			app.Users[i].Name = name
			app.Users[i].Email = email
			app.Users[i].PhoneNumber = phoneNumber
			break
		}
	}
}

// CreateTransaction allows users to send money to another user or a bank account
func (app *Application) CreateTransaction(transactionType string, fromUserID, toUserID int, amount float64, accountDetails ...string) int {
	transactionID := len(app.Transactions) + 1
	transaction := Transaction{
		ID:              transactionID,
		FromUserID:      fromUserID,
		ToUserID:        toUserID,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
	}
	app.Transactions = append(app.Transactions, transaction)
	return transactionID
}

// MakePayment allows users to make payment for the transaction via Card/UPI/Netbanking
func (app *Application) MakePayment(transactionID int, paymentMethod PaymentMethod, paymentDetails ...string) {
	// Implement payment logic here
	// For simplicity, we'll just print the payment details
	fmt.Printf("Payment successful for Transaction ID %d using %s\n", transactionID, paymentMethod)
}

// RefundTransaction allows users to refund a transaction
func (app *Application) RefundTransaction(transactionID int) {
	// Implement refund logic here
	// For simplicity, we'll just print the refund details
	fmt.Printf("Refund successful for Transaction ID %d\n", transactionID)
}

// ViewTransactionsHistory allows users to view transaction history
func (app *Application) ViewTransactionsHistory(userID int) []Transaction {
	var userTransactions []Transaction
	for _, transaction := range app.Transactions {
		if transaction.FromUserID == userID || transaction.ToUserID == userID {
			userTransactions = append(userTransactions, transaction)
		}
	}
	return userTransactions
}

func main() {
	// Create a new instance of the payment application
	app := Application{}

	// Register a new user
	userID := app.RegisterUser("1234567890", "password")

	// Update user profile
	app.UpdateUser(userID, "John Doe", "john@example.com", "1234567890")

	// Create a transaction to send money to another user
	transactionID := app.CreateTransaction("PAYTM", userID, 2, 50.0)

	// Make payment for the transaction using Card
	app.MakePayment(transactionID, Card, "CardDetails...")

	// Refund the transaction
	app.RefundTransaction(transactionID)

	// View transaction history
	transactions := app.ViewTransactionsHistory(userID)
	fmt.Println("Transaction History:")
	for _, transaction := range transactions {
		fmt.Printf("Transaction ID: %d, Type: %s, Amount: %.2f\n", transaction.ID, transaction.TransactionType, transaction.Amount)
	}
}
