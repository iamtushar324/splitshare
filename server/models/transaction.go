package models

import (
	"errors"
	"time"

	"github.com/iamtushar324/splitshare/server/db"
	"github.com/iamtushar324/splitshare/server/forms"
)

// Transaction represents a financial transaction between two users
type Transaction struct {
	ID             int64   `db:"id, primarykey, autoincrement" json:"id"`
	LenderID       int64   `db:"lender_id" json:"lender_id"`
	BorrowerID     int64   `db:"borrower_id" json:"borrower_id"`
	Amount         float64 `db:"amount" json:"amount"`
	TransactionDate time.Time `db:"transaction_date" json:"transaction_date"`
	IsSettled      bool    `db:"is_settled" json:"is_settled"`
	Description    string  `db:"description" json:"description,omitempty"`
	UpdatedAt      int64   `db:"updated_at" json:"-"`
	CreatedAt      int64   `db:"created_at" json:"-"`
}

// TransactionModel represents the model for handling transactions
type TransactionModel struct{}

// CreateTransaction creates a new transaction in the database
func (m TransactionModel) CreateTransaction(form forms.TransactionForm) (transaction Transaction, err error) {
	getDb := db.GetDB()

	err = getDb.QueryRow("INSERT INTO public.transactions (lender_id, borrower_id, amount, transaction_date, is_settled, description, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		form.LenderID, form.BorrowerID, form.Amount, form.TransactionDate, form.IsSettled, form.Description, time.Now().Unix(), time.Now().Unix()).Scan(&transaction.ID)
	if err != nil {
		return transaction, errors.New("something went wrong, please try again later")
	}

	transaction.LenderID = form.LenderID
	transaction.BorrowerID = form.BorrowerID
	transaction.Amount = form.Amount
	transaction.TransactionDate = form.TransactionDate
	transaction.IsSettled = form.IsSettled
	transaction.Description = form.Description
	transaction.CreatedAt = time.Now().Unix()
	transaction.UpdatedAt = time.Now().Unix()

	return transaction, err
}

// GetTransaction retrieves a transaction by its ID
func (m TransactionModel) GetTransaction(transactionID int64) (transaction Transaction, err error) {
	err = db.GetDB().SelectOne(&transaction, "SELECT id, lender_id, borrower_id, amount, transaction_date, is_settled, description, updated_at, created_at FROM public.transactions WHERE id=$1 LIMIT 1", transactionID)
	return transaction, err
}

// GetTransactionsByUser retrieves all transactions for a specific user (as lender or borrower)
func (m TransactionModel) GetTransactionsByUser(userID int64) (transactions []Transaction, err error) {
	_, err = db.GetDB().Select(&transactions, "SELECT id, lender_id, borrower_id, amount, transaction_date, is_settled, description, updated_at, created_at FROM public.transactions WHERE lender_id=$1 OR borrower_id=$2", userID, userID)
	return transactions, err
}

// SettleTransaction marks a transaction as settled
func (m TransactionModel) SettleTransaction(transactionID int64) (err error) {
	_, err = db.GetDB().Exec("UPDATE public.transactions SET is_settled=$1, updated_at=$2 WHERE id=$3", true, time.Now().Unix(), transactionID)
	return err
}
