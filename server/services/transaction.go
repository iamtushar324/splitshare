package services

import (
	"github.com/iamtushar324/splitshare/server/forms"
	"github.com/iamtushar324/splitshare/server/models"
)

type TransactionService struct {
	transactionModel models.TransactionModel
}

func (s *TransactionService) CreateTransaction(form forms.TransactionForm) (models.Transaction, error) {
	return s.transactionModel.CreateTransaction(form)
}

func (s *TransactionService) GetTransaction(transactionID int64) (models.Transaction, error) {
	return s.transactionModel.GetTransaction(transactionID)
}

func (s *TransactionService) GetTransactionsByUser(userID int64) ([]models.Transaction, error) {
	return s.transactionModel.GetTransactionsByUser(userID)
}

func (s *TransactionService) SettleTransaction(transactionID int64) error {
	return s.transactionModel.SettleTransaction(transactionID)
}
