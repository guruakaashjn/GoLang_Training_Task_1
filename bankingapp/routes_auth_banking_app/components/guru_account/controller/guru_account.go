package controller

import (
	accout_service "bankingapp/components/guru_account/service"
	customer_service "bankingapp/components/guru_customer/service"
	"bankingapp/guru_errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type InputTransaction struct {
	AccountNumber string
	Amount        int
	startDate     string
	endDate       string
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside create account controller function")
	var newAccountTemp *accout_service.Account
	err := json.NewDecoder(r.Body).Decode(&newAccountTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAccountError(guru_errors.CreateAccountFailed).GetSpecificMessage())
		// panic(err)
	}
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	// fmt.Println("Current Customer: ", currentCustomer)
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.CreateAccountFailed)
		panic(guru_errors.NewAccountError(guru_errors.CreateAccountFailed).GetSpecificMessage())
	}
	newAccount := currentCustomer.CreateAccount(newAccountTemp.BankId, newAccountTemp.Balance)
	// fmt.Println("New Account : ", newAccount)
	json.NewEncoder(w).Encode(newAccount)
	panic(guru_errors.NewAccountError(guru_errors.CreateAccountSuccess).GetSpecificMessage())

	// fmt.Println("Inside create account controller function POST REQUEST DONE")
}

func ReadAccountAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read all accounts controller function")
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
	}
	requiredAccounts := currentCustomer.ReadAllAccountsOfCustomer()
	json.NewEncoder(w).Encode(requiredAccounts)
	panic(guru_errors.NewAccountError(guru_errors.ReadAccountSuccess).GetSpecificMessage())
	// fmt.Println("Inside read all accounts controller function GET REQUEST DONE")

}

func ReadAccountById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read account by id controller function")
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
	}
	requiredAccount := currentCustomer.ReadAccountById(uuid.MustParse(slugs["account-id"]))
	if requiredAccount == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredAccount)
	panic(guru_errors.NewAccountError(guru_errors.ReadAccountSuccess).GetSpecificMessage())
	// fmt.Println("Inside read account by id controller function GET REQUEST DONE")
}

func UpdateAccountById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside update account by id controller function")
	json.NewEncoder(w).Encode(guru_errors.UpdateAccountSuccess)
	panic(guru_errors.NewAccountError(guru_errors.UpdateAccountSuccess).GetSpecificMessage())
	// fmt.Println("Inside update account by id controller function PUT REQUEST DONE")

}
func DeleteAccountById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside delete account by id controller function")
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))

	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.DeleteAccountFailed)
		panic(guru_errors.NewAccountError(guru_errors.DeleteAccountFailed).GetSpecificMessage())
	}

	deletedAccount := currentCustomer.DeleteAccount(uuid.MustParse(slugs["account-id"]))
	json.NewEncoder(w).Encode(deletedAccount)
	panic(guru_errors.NewAccountError(guru_errors.DeleteAccountSuccess).GetSpecificMessage())
	// fmt.Println("Inside delete account by id controller function DELETE REQUEST DONE")
}

func DepositMoney(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside deposit money controller function")
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.DepositFailed)
		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
	}
	updatedAccount := currentCustomer.DepositMoney(uuid.MustParse(slugs["account-id"]), newInputTransaction.Amount)
	json.NewEncoder(w).Encode(updatedAccount)
	panic(guru_errors.NewAccountError(guru_errors.DepositSuccess).GetSpecificMessage())

}
func WithdrawMoney(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside withdraw money controller function")
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.DepositFailed)
		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
	}
	updatedAccount := currentCustomer.WithdrawMoney(uuid.MustParse(slugs["account-id"]), newInputTransaction.Amount)
	json.NewEncoder(w).Encode(updatedAccount)

	panic(guru_errors.NewAccountError(guru_errors.WithdrawSuccess).GetSpecificMessage())
}
func TransferMoney(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside transfer money controller function")
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.TransferFailed)
		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
	}

	updatedSenderAccount, updatedReceiverAccount := currentCustomer.TransferMoney(uuid.MustParse(slugs["account-id"]), uuid.MustParse(newInputTransaction.AccountNumber), newInputTransaction.Amount)

	tempArray := [2]*accout_service.Account{updatedSenderAccount, updatedReceiverAccount}
	json.NewEncoder(w).Encode(tempArray)

	panic(guru_errors.NewAccountError(guru_errors.TransferSuccess).GetSpecificMessage())
}

func PassbookPrint(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside passbook print controller function")
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewAccountError(guru_errors.PassbookFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
	if currentCustomer == nil {
		json.NewEncoder(w).Encode(guru_errors.TransferFailed)
		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
	}
	requiredPassbookInRange := currentCustomer.GetPassbookInRange(uuid.MustParse(newInputTransaction.AccountNumber), newInputTransaction.startDate, newInputTransaction.endDate)

	json.NewEncoder(w).Encode(requiredPassbookInRange)
	panic(guru_errors.NewAccountError(guru_errors.PassbookSuccess).GetSpecificMessage())

}
