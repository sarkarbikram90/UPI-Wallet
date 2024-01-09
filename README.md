# Go Wallet
# Introduction:

Go Wallet is a command-line payment application written in Go (Golang). It provides users with a simple yet powerful interface to manage their transactions, send money to other users, and view transaction history.

# Features:

# Account Management:

Users can create accounts by providing their phone number and password.
Profile details, including name, email, and phone number, can be updated.
Transaction Handling:

Users can send money to other users within the application.
Transactions support two types: PAYTM (between users) and BANK (to a bank account).
Payments can be made using different methods such as Card, UPI, and Netbanking.
Refund and Transaction History:

Users have the option to refund a transaction, and the refunded amount goes to the original source.
The application provides a feature to view transaction history.
Code Structure:

The project is organized into packages, each serving a specific purpose:

main.go: The entry point of the application.
app: Contains modules related to the application, such as user management, transactions, and payments.
test: Includes test files for the application.
Usage:

Register User:

# RegisterUser [phone_number] [password]
Update User Profile:

# UpdateUser [user_id] [name] [email] [phone_number]
Create Transaction (PAYTM):

# CreateTransaction PAYTM [from_user_id] [to_user_id] [amount]
Create Transaction (BANK):

# ViewTransactionsHistory [user_id]

# CreateTransaction BANK [from_user_id] [account_number] [ifsc_code] [amount]
Make Payment:

# MakePayment [transaction_id] [payment_method] [... details related to payment method ...]
Refund Transaction:

# RefundTransaction [transaction_id]
View Transaction History:

# ViewTransactionsHistory [user_id]
How to Run:

# Clone the repository: git clone https://github.com/sarkarbikram90/go-wallet.git
Navigate to the project directory: # cd go-wallet
Run the application: # go run main.go

Contributions to Go Wallet are welcome! Feel free to open issues, propose enhancements, or submit pull requests.

License:

This project is licensed under the MIT License - see the LICENSE file for details.
