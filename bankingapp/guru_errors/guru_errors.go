package guru_errors

import "errors"

type AdminError struct {
	errorType       error
	specificMessage string
}

const NotAdmin string = "user not an admin"
const CreatedAdmin string = "admin created successfully"
const DeletedAdmin string = "admin is deleted"

func NewAdminError(specificMessage string) *AdminError {
	return &AdminError{
		errorType:       errors.New("previlage error"),
		specificMessage: specificMessage,
	}
}

func (e *AdminError) GetSpecificMessage() string {
	return e.specificMessage
}

type UserError struct {
	errorType       error
	specificMessage string
}

const DeletedUser string = "user is deleted successfully"
const NotExistUser string = "user doesn't exist"
const DeletedUserStatus string = "user is deleted"
const UpdatedUser string = "user is updated successfully"
const CreatedUser string = "user is created successfully"
const ReadUser string = "user read done successfully"

func NewUserError(specificMessage string) *UserError {
	return &UserError{
		errorType:       errors.New("user error"),
		specificMessage: specificMessage,
	}
}
func (e *UserError) GetSpecificMessage() string {
	return e.specificMessage
}

type AccountError struct {
	errorType       error
	specificMessage string
}

const DeletedAccount string = "account deleted successfully"
const UpdatedAccount string = "account updated successfully"
const CreatedAccount string = "account created successfully"
const ReadAccount string = "account read done successfully"

const DeletedAccountStatus string = "account is deleted"
const DeletedAccountAlready string = "account is deleted already"
const NotExistAccount string = "account doesn't exist"
const InSufficientBalance string = "account balance in-sufficient"

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
const DeletedBank string = "bank deleted successfully"
const UpdatedBank string = "bank details updated successfully"
const CreatedBank string = "bank created successfully"
const ReadBank string = "bank read done successfully"
const DeletedBankStatus string = "bank is deleted"

func NewBankError(specificMessage string) *BankError {
	return &BankError{
		errorType:       errors.New("account error"),
		specificMessage: specificMessage,
	}
}
func (e *BankError) GetSpecificMessage() string {
	return e.specificMessage
}
