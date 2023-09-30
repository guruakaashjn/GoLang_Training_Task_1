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

// SERVICE ERRORS
const DeletedUser string = "user is deleted successfully"
const NotExistUser string = "user doesn't exist"
const DeletedUserStatus string = "user is deleted"
const UpdatedUser string = "user is updated successfully"
const CreatedUser string = "user is created successfully"
const ReadUser string = "user read done successfully"
const UserTotalBalance string = "user total balance calculated successfully"
const UserAccoutBalanceMap string = "user account balance corresponding to each account listed cussessfully"
const BankNetWorthMap string = "total bank net worth corresponding to each bank listed successfully"

// CONTROLLER ERRORS
const CreateAdminFailed string = "create admin failed from customer controller"
const CreateCustomerFailed string = "create customer failed from customer controller"
const CreateAdminSuccess string = "create admin success from customer controller"
const CreateCustomerSuccess string = "create customer success from customer controller"
const ReadCustomerFailed string = "read customer failed from customer controller"
const ReadCustomerSuccess string = "read customer success from customer controller"
const UpdateCustomerFailed string = "update customer failed from customer controller"
const UpdateCustomerSuccess string = "update customer success from customer controller"
const DeleteCustomerFailed string = "delete customer failed from customer controller"
const DeleteCustomerSuccess string = "delete customer success from customer controller"

const TotalBalanceFailed string = "total balance of customer failed from customer controller"
const TotalBalanceSuccess string = "total balance of customer success from customer controller"
const AccountBalanceListFailed string = "account balance list of customer failed from customer controller"
const AccountBalanceListSuccess string = "account balance list of customer success from customer controller"

const AdminObjectNotFound string = "admin object not found"
const AdminObjectFound string = "admin object found successfully"

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

// SERVICE ERRORS
const DeletedAccount string = "account deleted successfully"
const UpdatedAccount string = "account updated successfully"
const CreatedAccount string = "account created successfully"
const ReadAccount string = "account read done successfully"

const DeletedAccountStatus string = "account is deleted"
const DeletedAccountAlready string = "account is deleted already"
const NotExistAccount string = "account doesn't exist"
const InSufficientBalance string = "account balance in-sufficient"

const MoneyDeposited string = "money deposited successfully"
const MoneyDepositedError string = "money deposited un-successful"
const MoneyWithdraw string = "money withdraw successfully"
const MoneyWithdrawError string = "money withdraw un-successful"
const MoneyTransfered string = "money transfered successfully"
const MoneyTransferedError string = "money transfered un-successful"

const PassbookReadInRange string = "passbook read in given range done successfully"

// CONTOLLER ERRORS
const CreateAccountFailed string = "create account failed from account controller"
const CreateAccountSuccess string = "create account success from account controller"
const ReadAccountFailed string = "read account failed from account controller"
const ReadAccountSuccess string = "read account success from account controller"
const UpdateAccountFailed string = "update account failed from account controller"
const UpdateAccountSuccess string = "update account success from account controller"
const DeleteAccountFailed string = "delete account failed from account controller"
const DeleteAccountSuccess string = "delete account success from account controller"

const DepositFailed string = "deposit of money to the account failed from account controller"
const DepositSuccess string = "deposit of money to the account success from account controller"
const WithdrawFailed string = "withdraw of money to the account failed from account controller"
const WithdrawSuccess string = "withdraw of money to the account success from account controller"
const TransferFailed string = "transfer of money to the account failed from account controller"
const TransferSuccess string = "transfer of money to the account success from account controller"

const PassbookFailed string = "passbook in range failed from account controller"
const PassbookSuccess string = "passbook in range success from account controller"

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

// SERVICE ERRORS
const BankContainsAccounts string = "bank cannot be deleted as it contains open bank accounts"
const DeletedBank string = "bank deleted successfully"
const UpdatedBank string = "bank details updated successfully"
const CreatedBank string = "bank created successfully"
const ReadBank string = "bank read done successfully"
const DeletedBankStatus string = "bank is deleted"
const ReadBankTransferAllMap string = "bank transfer all map read done successfully"

// CONTROLLER ERRORS
const CreateBankFailed string = "create bank failed from bank controller"
const CreateBankSuccess string = "create bank success from bank controller"
const ReadBankFailed string = "read bank failed from bank controller"
const ReadBankSuccess string = "read bank success from bank controller"
const UpdateBankFailed string = "update bank failed from bank controller"
const UpdateBankSuccess string = "update bank success from bank controller"
const DeleteBankFailed string = "delete bank failed from bank controller"
const DeleteBankSuccess string = "delete bank success from bank controller"

const NetWorthEachBankFailed string = "networth of each bank failed from bank controller"
const NetWorthEachBankSuccess string = "networth of each bank success from bank controller"
const NetWorthGivenBankFailed string = "networth of given bank failed from bank controller"
const NetWorthGivenBankSuccess string = "networth of given bank success from bank controller"

const BankNameBalanceMapFailed string = "bank name balance corresponding to each bank failed from bank controller"
const BankNameBalanceMapSuccess string = "bank name balance corresponding to each bank success from bank controller"
const BankNameBalanceMapAllFailed string = "all bank name balance corresponding to each bank failed from bank controller"
const BankNameBalanceMapAllSuccess string = "all bank name balance corresponding to each bank success from bank controller"

func NewBankError(specificMessage string) *BankError {
	return &BankError{
		errorType:       errors.New("account error"),
		specificMessage: specificMessage,
	}
}
func (e *BankError) GetSpecificMessage() string {
	return e.specificMessage
}

type authenticationError struct {
	Error           error
	specificMessage string
}

const AuthenticationFailed string = "login failed authentication failed"
const AuthenticationSuccess string = "login success authentication success"

func NewAuthenticationError(specificMessage string) *authenticationError {
	return &authenticationError{
		Error:           errors.New("authentication error"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *authenticationError) GetSpecificMessage() string {
	return e.specificMessage
}
