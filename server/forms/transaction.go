package forms

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionForm struct {
	LenderID       int64     `form:"lender_id" json:"lender_id" binding:"required"`
	BorrowerID     int64     `form:"borrower_id" json:"borrower_id" binding:"required"`
	Amount         float64   `form:"amount" json:"amount" binding:"required,gt=0"`
	TransactionDate time.Time `form:"transaction_date" json:"transaction_date" binding:"required"`
	IsSettled      bool      `form:"is_settled" json:"is_settled"`
	Description    string    `form:"description" json:"description" binding:"max=255"`
}

type GetTransactionByIdForm struct {
	TransactionId int64 `form:"transaction_id" json:"transaction_id" binding:"required"`
}

type TransactionFormValidator struct{}

func (f TransactionFormValidator) LenderID(tag string) string {
	switch tag {
	case "required":
		return "Lender ID is required"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f TransactionFormValidator) BorrowerID(tag string) string {
	switch tag {
	case "required":
		return "Borrower ID is required"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f TransactionFormValidator) Amount(tag string) string {
	switch tag {
	case "required":
		return "Amount is required"
	case "gt":
		return "Amount must be greater than zero"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f TransactionFormValidator) TransactionDate(tag string) string {
	switch tag {
	case "required":
		return "Transaction date is required"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f TransactionFormValidator) Description(tag string) string {
	switch tag {
	case "max":
		return "Description must be less than 255 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f TransactionFormValidator) ValidateTransaction(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "LenderID":
				return f.LenderID(err.Tag())
			case "BorrowerID":
				return f.BorrowerID(err.Tag())
			case "Amount":
				return f.Amount(err.Tag())
			case "TransactionDate":
				return f.TransactionDate(err.Tag())
			case "Description":
				return f.Description(err.Tag())
			}
		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

func (f TransactionFormValidator) ValidateGetTransactionFromIdParams(params map[string]string) error {
	var validate *validator.Validate
	validate = validator.New()

	for key, value := range params {
		switch key {
		case "transactionId":
			if err := validate.Var(value, "required,number"); err != nil {
				return err
			}
		default:
			return errors.New("invalid query parameter: " + key)
		}
	}

	return nil
}
