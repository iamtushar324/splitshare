package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamtushar324/splitshare/server/forms"
	"github.com/iamtushar324/splitshare/server/services"
)

type TransactionController struct {
	transactionService services.TransactionService
	transactionFormValidator forms.TransactionFormValidator
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		transactionService: services.TransactionService{},
		transactionFormValidator: forms.TransactionFormValidator{},
	}
}

func (ctrl TransactionController) CreateTransaction(c *gin.Context) {
	var transactionForm forms.TransactionForm

	if validationErr := c.ShouldBindJSON(&transactionForm); validationErr != nil {
		message := ctrl.transactionFormValidator.ValidateTransaction(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	transaction, err := ctrl.transactionService.CreateTransaction(transactionForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}

func (ctrl TransactionController) GetTransaction(c *gin.Context) {
	transactionIdString := c.Param("id")
	transactionId , err:= strconv.ParseInt(transactionIdString, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction_id"})
		return
	}

	transaction, err := ctrl.transactionService.GetTransaction(transactionId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (ctrl TransactionController) GetTransactionsByUser(c *gin.Context) {
	userID := getUserID(c)
	transactions, err := ctrl.transactionService.GetTransactionsByUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (ctrl TransactionController) SettleTransaction(c *gin.Context) {
	transactionIdString := c.Param("id")
	transactionId , err:= strconv.ParseInt(transactionIdString, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction_id"})
		return
	}

	err = ctrl.transactionService.SettleTransaction(transactionId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction settled successfully"})
}
