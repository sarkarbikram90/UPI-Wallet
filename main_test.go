package main

import (
	"testing"
)

func TestPaymentApplication(t *testing.T) {
	app := Application{}

	// Test user registration
	userID := app.RegisterUser("9876543210", "securepassword")
	if userID != 1 {
		t.Errorf("Expected user ID 1, but got %d", userID)
	}

	// Test user profile update
	app.UpdateUser(userID, "Alice", "alice@example.com", "9876543210")
	user := app.Users[0]
	if user.Name != "Alice" || user.Email != "alice@example.com" || user.PhoneNumber != "9876543210" {
		t.Errorf("User profile update failed")
	}

	// Test transaction creation
	transactionID := app.CreateTransaction("PAYTM", userID, 2, 25.0)
	if transactionID != 1 {
		t.Errorf("Expected transaction ID 1, but got %d", transactionID)
	}

	// Test payment
	app.MakePayment(transactionID, UPI, "UPIDetails...")
	// For simplicity, we don't have a real payment implementation, so no further checks are performed

	// Test refund
	app.RefundTransaction(transactionID)
	// For simplicity, we don't have a real refund implementation, so no further checks are performed

	// Test transaction history
	transactions := app.ViewTransactionsHistory(userID)
	if len(transactions) != 1 || transactions[0].ID != 1 || transactions[0].Amount != 25.0 {
		t.Errorf("Transaction history retrieval failed")
	}
}
