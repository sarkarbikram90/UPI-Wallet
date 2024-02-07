package main

import (
	"fmt"
	"strings"
)

// User represents a user of the application
type User struct {
	ID                string // Unique user ID generated based on name and mobile number
	Name              string
	PhoneNumber       string
	BankAccountNumber string
	AccountMoney      float64 // Balance
	PIN               string  // 4 digit UPI PIN
}

// Transaction represents a payment transaction
type Transaction struct {
	FromUserID string
	ToUserID   string
	Amount     float64
}

// Application represents the payment application
type Application struct {
	Users        []User
	Transactions []Transaction
}

// RegisterUser allows users to create accounts via their phone number, bank account number, and PIN
func (app *Application) RegisterUser() {
	var name, phoneNumber, bankAccountNumber, pin string

	fmt.Println("Welcome to the UPI payment wallet, to continue please register")
	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Please enter your mobile number: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Please enter your bank account number: ")
	fmt.Scanln(&bankAccountNumber)

	fmt.Print("Please create your UPI PIN: ")
	fmt.Scanln(&pin)

	// Generate unique user ID based on name and mobile number
	userID := generateUserID(name, phoneNumber)

	user := User{
		ID:                userID,
		PhoneNumber:       phoneNumber,
		PIN:               pin,
		Name:              name,
		BankAccountNumber: bankAccountNumber,
	}
	app.Users = append(app.Users, user)
	fmt.Printf("Welcome to your UPI wallet account, %s!\n", user.Name)
	fmt.Printf("Your user ID is: %s\n", user.ID)
}

// generateUserID generates a unique user ID based on name and mobile number
func generateUserID(name, phoneNumber string) string {
	// Remove spaces from the name and concatenate with the mobile number
	userID := strings.ReplaceAll(name, " ", "") + "-" + phoneNumber
	return userID
}

// SendMoney allows users to send money to another user
func (app *Application) SendMoney() {
	var senderID, recipientID string
	var amount float64

	fmt.Print("Enter your user ID: ")
	fmt.Scanln(&senderID)

	fmt.Print("Enter recipient's user ID: ")
	fmt.Scanln(&recipientID)

	fmt.Print("Enter amount to send: ")
	fmt.Scanln(&amount)

	// Find the sender and recipient by ID
	var senderIndex, recipientIndex int
	for i, user := range app.Users {
		if user.ID == senderID {
			senderIndex = i
		}
		if user.ID == recipientID {
			recipientIndex = i
		}
	}

	if senderIndex == -1 || recipientIndex == -1 {
		fmt.Println("Sender or recipient not found.")
		return
	}

	// Check if sender has sufficient balance
	if app.Users[senderIndex].AccountMoney < amount {
		fmt.Println("Insufficient balance.")
		return
	}

	// Perform the transaction
	app.Users[senderIndex].AccountMoney -= amount
	app.Users[recipientIndex].AccountMoney += amount

	fmt.Printf("Payment successful. New balance of sender (%s): %.2f, recipient (%s): %.2f\n",
		app.Users[senderIndex].Name, app.Users[senderIndex].AccountMoney,
		app.Users[recipientIndex].Name, app.Users[recipientIndex].AccountMoney)
}

// AddMoney allows users to add money to their wallet
func (app *Application) AddMoney() {
	var userID string
	var amount float64

	fmt.Print("Please enter your user ID: ")
	fmt.Scanln(&userID)

	fmt.Print("Enter amount to add: ")
	fmt.Scanln(&amount)

	// Find the user by ID
	var user *User
	for i := range app.Users {
		if app.Users[i].ID == userID {
			user = &app.Users[i]
			break
		}
	}

	if user == nil {
		fmt.Println("User not found.")
		return
	}

	// Update user's account balance
	user.AccountMoney += amount
	fmt.Printf("Amount has been added, your account balance is %.2f!\n", user.AccountMoney)
}

// CheckBalance allows users to check their account balance
func (app *Application) CheckBalance() {
	var userID string

	fmt.Print("Enter your user ID: ")
	fmt.Scanln(&userID)

	// Find the user by ID
	for _, user := range app.Users {
		if user.ID == userID {
			fmt.Printf("Your current balance: %.2f\n", user.AccountMoney)
			return
		}
	}

	fmt.Println("User not found.")
}

func main() {
	// Create a new instance of the payment application
	app := Application{}

	// Continue loop
	for {
		fmt.Println("\nWelcome to UPI Wallet, to continue choose an option:")
		fmt.Println("1. Register")
		fmt.Println("2. Add money")
		fmt.Println("3. Send money")
		fmt.Println("4. Check balance")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			app.RegisterUser()
		case 2:
			app.AddMoney()
		case 3:
			app.SendMoney()
		case 4:
			app.CheckBalance()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		var continueOption string
		fmt.Print("\nDo you want to perform any other operations? (yes/no): ")
		fmt.Scanln(&continueOption)
		if continueOption == "no" {
			fmt.Println("Thank you for choosing UPI Wallet")
			break
		}
	}
}
