package controller

import (
	"bankingapp/errors"
	"bankingapp/models/account"
	"bankingapp/models/entry"
	"bankingapp/models/offer"
	"bankingapp/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InputTransaction struct {
	AccountNumber uint
	Amount        int
}

func (controller *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := account.Account{}
	err := web.UnmarshalJSON(r, &newAccount)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	slugs := mux.Vars(r)
	idTemp, _ := strconv.Atoi(slugs["customer-id"])

	newAccount.CustomerID = uint(idTemp)
	err = controller.service.CreateAccount(&newAccount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newAccount)
}

func (controller *AccountController) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	limit, offset := web.ParseLimitAndOffset(r)
	givenAssociations := web.ParsePreloading(r)

	allAccounts := &[]account.Account{}
	var totalCount int
	err = controller.service.GetAllAccounts(allAccounts, &totalCount, uint(idTemp), limit, offset, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allAccounts)
}

func (controller *AccountController) GetAccountById(w http.ResponseWriter, r *http.Request) {
	requiredAccount := account.Account{}
	slugs := mux.Vars(r)
	customerIdTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	givenAssociations := web.ParsePreloading(r)
	requiredAccount.ID = uint(idTemp)
	requiredAccount.CustomerID = uint(customerIdTemp)

	err = controller.service.GetAccountById(&requiredAccount, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredAccount)
}
func (controller *AccountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("account to update")
	accountToUpdate := account.Account{}
	fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &accountToUpdate)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	accountToUpdate.ID = uint(intId)
	intId, _ = strconv.Atoi(slugs["customer-id"])

	accountToUpdate.CustomerID = uint(intId)
	fmt.Println("account to update")
	fmt.Println(&accountToUpdate)
	err = controller.service.UpdateAccount(&accountToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, accountToUpdate)
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	controller.log.Print("delete account call")
	accountToDelete := account.Account{}
	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	accountToDelete.ID = uint(intID)
	err = controller.service.DeleteAccount(&accountToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete account successful.")
}

func (controller *AccountController) TransferMoney(w http.ResponseWriter, r *http.Request) {

	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	senderAccount := account.Account{}
	senderAccount.ID = uint(intID)

	inputTransaction := InputTransaction{}
	fmt.Println(r.Body)
	err = web.UnmarshalJSON(r, &inputTransaction)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	receiverAccount := account.Account{}
	receiverAccount.ID = inputTransaction.AccountNumber
	amount := inputTransaction.Amount

	err = controller.service.TransferMoney(&senderAccount, &receiverAccount, amount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Transfer Money done successful.")
}

func (controller *AccountController) DepositMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------------------------------------------FIRST-----------------------------------")
	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	fmt.Println("--------------------------------------------------1-----------------------------------")
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	senderAccount := account.Account{}
	senderAccount.ID = uint(intID)

	inputTransaction := InputTransaction{}
	fmt.Println("--------------------------------------------------2-----------------------------------")
	fmt.Println(r.Body)
	err = web.UnmarshalJSON(r, &inputTransaction)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	amount := inputTransaction.Amount

	err = controller.service.DepositMoney(&senderAccount, amount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Deposit Money done successful.")
}
func (controller *AccountController) WithdrawMoney(w http.ResponseWriter, r *http.Request) {

	var err error
	slugs := mux.Vars(r)
	intID, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	senderAccount := account.Account{}
	senderAccount.ID = uint(intID)

	inputTransaction := InputTransaction{}
	fmt.Println(r.Body)
	err = web.UnmarshalJSON(r, &inputTransaction)
	if err != nil {
		fmt.Println("error from UnMarshalJSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	amount := inputTransaction.Amount

	err = controller.service.WithdrawMoney(&senderAccount, amount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Withdraw Money done successful.")
}

func (controller *AccountController) PassbookPrint(w http.ResponseWriter, r *http.Request) {

	requiredAccount := &account.Account{}
	limit, offset := web.ParseLimitAndOffset(r)
	var totalCount int
	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	requiredAccount.ID = uint(idTemp)

	startDateTimeDotTime, endDateTimeDotTime := web.ParseStartDateEndDate(r)

	// fmt.Println(startDate, "    ", endDate)
	requiredEntries := &[]entry.Entry{}
	err = controller.service.PassbookPrint(requiredAccount, startDateTimeDotTime, endDateTimeDotTime, requiredEntries, limit, offset, &totalCount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, requiredEntries)

}

func (controller *AccountController) GetAvailableOffers(w http.ResponseWriter, r *http.Request) {

	slugs := mux.Vars(r)
	customerIdTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	givenAssociations := web.ParsePreloading(r)

	availableOffers := &[]offer.Offer{}
	requiredAccount := &account.Account{}
	requiredAccount.ID = uint(idTemp)
	requiredAccount.CustomerID = uint(customerIdTemp)

	err = controller.service.GetAvailableOffers(requiredAccount, availableOffers, givenAssociations)

	// err = controller.service.GetAccountById(&requiredAccount, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, availableOffers)
}

func (controller *AccountController) TakeAvailableOffer(w http.ResponseWriter, r *http.Request) {
	slugs := mux.Vars(r)
	customerIdTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	takeAvailableOffer := &offer.Offer{}

	fmt.Println(r.Body)

	err = web.UnmarshalJSON(r, &takeAvailableOffer)
	if err != nil {
		fmt.Println("error from unmarshal JSON")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	requiredAccount := &account.Account{}
	requiredAccount.ID = uint(idTemp)
	requiredAccount.CustomerID = uint(customerIdTemp)

	err = controller.service.TakeAvailableOffer(requiredAccount, takeAvailableOffer)

	// err = controller.service.GetAccountById(&requiredAccount, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, takeAvailableOffer)
}

func (controller *AccountController) ChosenOffers(w http.ResponseWriter, r *http.Request) {
	slugs := mux.Vars(r)
	customerIdTemp, err := strconv.Atoi(slugs["customer-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	chosenOffers := &[]offer.Offer{}

	requiredAccount := &account.Account{}
	requiredAccount.ID = uint(idTemp)
	requiredAccount.CustomerID = uint(customerIdTemp)

	err = controller.service.ChosenOffers(requiredAccount, chosenOffers)

	// err = controller.service.GetAccountById(&requiredAccount, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, chosenOffers)
}

// package controller

// import (
// 	accout_service "bankingapp/components/guru_account/service"
// 	customer_service "bankingapp/components/guru_customer/service"
// 	"bankingapp/guru_errors"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// type InputTransaction struct {
// 	AccountNumber string
// 	Amount        int
// 	StartDate     string
// 	EndDate       string
// }

// func CreateAccount(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside create account controller function")
// 	var newAccountTemp *accout_service.Account
// 	err := json.NewDecoder(r.Body).Decode(&newAccountTemp)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAccountError(guru_errors.CreateAccountFailed).GetSpecificMessage())
// 		// panic(err)
// 	}
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	// fmt.Println("Current Customer: ", currentCustomer)
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.CreateAccountFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.CreateAccountFailed).GetSpecificMessage())
// 	}
// 	newAccount := currentCustomer.CreateAccount(newAccountTemp.BankId, newAccountTemp.Balance)
// 	// fmt.Println("New Account : ", newAccount)
// 	json.NewEncoder(w).Encode(newAccount)
// 	panic(guru_errors.NewAccountError(guru_errors.CreateAccountSuccess).GetSpecificMessage())

// 	// fmt.Println("Inside create account controller function POST REQUEST DONE")
// }

// func ReadAccountAll(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside read all accounts controller function")
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
// 	}
// 	requiredAccounts := currentCustomer.ReadAllAccountsOfCustomer()
// 	json.NewEncoder(w).Encode(requiredAccounts)
// 	panic(guru_errors.NewAccountError(guru_errors.ReadAccountSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside read all accounts controller function GET REQUEST DONE")

// }

// func ReadAccountById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside read account by id controller function")
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
// 	}
// 	requiredAccount := currentCustomer.ReadAccountById(uuid.MustParse(slugs["account-id"]))
// 	if requiredAccount == nil {
// 		json.NewEncoder(w).Encode(guru_errors.ReadAccountFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.ReadAccountFailed).GetSpecificMessage())
// 	}
// 	json.NewEncoder(w).Encode(requiredAccount)
// 	panic(guru_errors.NewAccountError(guru_errors.ReadAccountSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside read account by id controller function GET REQUEST DONE")
// }

// func UpdateAccountById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside update account by id controller function")
// 	json.NewEncoder(w).Encode(guru_errors.UpdateAccountSuccess)
// 	panic(guru_errors.NewAccountError(guru_errors.UpdateAccountSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside update account by id controller function PUT REQUEST DONE")

// }
// func DeleteAccountById(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside delete account by id controller function")
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))

// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DeleteAccountFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.DeleteAccountFailed).GetSpecificMessage())
// 	}

// 	deletedAccount := currentCustomer.DeleteAccount(uuid.MustParse(slugs["account-id"]))
// 	json.NewEncoder(w).Encode(deletedAccount)
// 	panic(guru_errors.NewAccountError(guru_errors.DeleteAccountSuccess).GetSpecificMessage())
// 	// fmt.Println("Inside delete account by id controller function DELETE REQUEST DONE")
// }

// func DepositMoney(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside deposit money controller function")
// 	var newInputTransaction *InputTransaction
// 	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
// 	}
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DepositFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
// 	}
// 	updatedAccount := currentCustomer.DepositMoney(uuid.MustParse(slugs["account-id"]), newInputTransaction.Amount)
// 	json.NewEncoder(w).Encode(updatedAccount)
// 	panic(guru_errors.NewAccountError(guru_errors.DepositSuccess).GetSpecificMessage())

// }
// func WithdrawMoney(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside withdraw money controller function")
// 	var newInputTransaction *InputTransaction
// 	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
// 	}
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.DepositFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.DepositFailed).GetSpecificMessage())
// 	}
// 	updatedAccount := currentCustomer.WithdrawMoney(uuid.MustParse(slugs["account-id"]), newInputTransaction.Amount)
// 	json.NewEncoder(w).Encode(updatedAccount)

// 	panic(guru_errors.NewAccountError(guru_errors.WithdrawSuccess).GetSpecificMessage())
// }
// func TransferMoney(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside transfer money controller function")
// 	var newInputTransaction *InputTransaction
// 	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
// 	}
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.TransferFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
// 	}

// 	updatedSenderAccount, updatedReceiverAccount := currentCustomer.TransferMoney(uuid.MustParse(slugs["account-id"]), uuid.MustParse(newInputTransaction.AccountNumber), newInputTransaction.Amount)

// 	tempArray := [2]*accout_service.Account{updatedSenderAccount, updatedReceiverAccount}
// 	json.NewEncoder(w).Encode(tempArray)

// 	panic(guru_errors.NewAccountError(guru_errors.TransferSuccess).GetSpecificMessage())
// }

// func PassbookPrint(w http.ResponseWriter, r *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	fmt.Println("Inside passbook print controller function")
// 	var newInputTransaction *InputTransaction
// 	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 		panic(guru_errors.NewAccountError(guru_errors.PassbookFailed).GetSpecificMessage())
// 	}
// 	slugs := mux.Vars(r)
// 	currentCustomer := customer_service.ReadCustomerById(uuid.MustParse(slugs["customer-id"]))
// 	if currentCustomer == nil {
// 		json.NewEncoder(w).Encode(guru_errors.TransferFailed)
// 		panic(guru_errors.NewAccountError(guru_errors.TransferFailed).GetSpecificMessage())
// 	}
// 	// fmt.Println("Passbook PRint Func : ", newInputTransaction.StartDate, "    ", newInputTransaction.EndDate)
// 	requiredPassbookInRange := currentCustomer.GetPassbookInRange(uuid.MustParse(slugs["account-id"]), newInputTransaction.StartDate, newInputTransaction.EndDate)

// 	json.NewEncoder(w).Encode(requiredPassbookInRange)
// 	panic(guru_errors.NewAccountError(guru_errors.PassbookSuccess).GetSpecificMessage())

// }
