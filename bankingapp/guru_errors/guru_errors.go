package guru_errors

import "errors"

type AdminError struct {
	errorType       error
	specificMessage string
}

const NotAdmin string = "user not an admin"

func NewAdminError(specificMessage string) *AdminError {
	return &AdminError{
		errorType:       errors.New("previlage error"),
		specificMessage: specificMessage,
	}
}

func (e *AdminError) GetSpecificMessage() string {
	return e.specificMessage
}

type NotAUser struct {
	errorType       error
	specificMessage string
}

const DeletedUser string = "user is deleted"
const NotExistUser string = "user doesn't exist"
const UpdatedUser string = "user is updated"

func NewNotAUser(specificMessage string) *NotAUser {
	return &NotAUser{
		errorType:       errors.New("user error"),
		specificMessage: specificMessage,
	}
}
func (e *NotAUser) GetSpecificMessage() string {
	return e.specificMessage
}

type AccountError struct {
	errorType       error
	specificMessage string
}

const DeletedAccount string = "account deleted"
const UpdatedAccount string = "account updated"
const DeletedAccountStatus string = "account is deleted"
const DeletedAccountAlready string = "account is deleted already"
const NotExistAccount string = "account doesn't exist"
const InSufficientBalance string = "account initial balance in-sufficient"

func NewAccountError(specificMessage string) *AccountError {
	return &AccountError{
		errorType:       errors.New("account error"),
		specificMessage: specificMessage,
	}
}
func (e *AccountError) GetSpecificMessage() string {
	return e.specificMessage
}

type BankError struct {
	errorType       error
	specificMessage string
}

const BankContainsAccounts string = "bank cannot be deleted as it contains open bank accounts"
const DeletedBank string = "bank deleted"
const UpdatedBank string = "bank details updated"

func NewBankError(specificMessage string) *BankError {
	return &BankError{
		errorType:       errors.New("account error"),
		specificMessage: specificMessage,
	}
}
func (e *BankError) GetSpecificMessage() string {
	return e.specificMessage
}
