package guru_bank

import (
	bank_service "bankingapp/components/guru_bank/service"
	customer_service "bankingapp/components/guru_customer/service"
	"bankingapp/guru_errors"
	"bankingapp/middleware/auth"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// var initialAdmin *customer_service.Customer = &customer_service.Customer{
// 	CustomerId:   uuid.MustParse("7affab7a-5c59-11ee-8c99-0242ac120002"),
// 	FirstName:    "Admin Initial",
// 	LastName:     "Admin surname",
// 	TotalBalance: 0,
// 	IsAdmin:      true,
// 	IsActive:     true,
// 	Accounts:     make([]*account_service.Account, 0),
// }

type InputTransaction struct {
	BankId    string
	StartDate string
	EndDate   string
}

func GetAdminObjectFromCookie(w http.ResponseWriter, r *http.Request) (requiredAdmin *customer_service.Customer) {
	defer func() {
		if details := recover(); details != nil {
			fmt.Println(details)
		}
	}()
	cookie, err1 := r.Cookie("authone")
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err1)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}
	token := cookie.Value
	payload, err2 := auth.Verify(token)
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err2)
		panic(guru_errors.NewUserError(guru_errors.AdminObjectNotFound).GetSpecificMessage())
	}
	requiredAdminTemp := customer_service.ReadCustomerByUserName(payload.UserName)
	if requiredAdminTemp == nil {
		json.NewEncoder(w).Encode(guru_errors.DeletedAdmin)
	}
	requiredAdmin = requiredAdminTemp
	panic(guru_errors.NewUserError(guru_errors.AdminObjectFound).GetSpecificMessage())

}

func CreateBank(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside create bank controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var newBankTemp *bank_service.Bank
	err := json.NewDecoder(r.Body).Decode(&newBankTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewBankError(guru_errors.CreateBankFailed).GetSpecificMessage())
	}

	newBank := adminObject.CreateBank(newBankTemp.GetBankName())
	if newBank == nil {
		json.NewEncoder(w).Encode(guru_errors.CreateBankFailed)
		panic(guru_errors.NewBankError(guru_errors.CreateBankFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(newBank)
	// fmt.Println("Inside create bank controller function POST REQUEST DONE")
	panic(guru_errors.NewBankError(guru_errors.CreateBankSuccess).GetSpecificMessage())
}

func ReadBankAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read bank all controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	requiredBanks := adminObject.ReadAllBanks()
	if requiredBanks == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadBankFailed)
		panic(guru_errors.NewBankError(guru_errors.ReadBankFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredBanks)
	// fmt.Println("Inside read bank all controller function GET REQUEST DONE")
	panic(guru_errors.NewBankError(guru_errors.ReadBankSuccess).GetSpecificMessage())
}

func ReadBankById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside read bank by id controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	requiredBank := adminObject.ReadBankById(uuid.MustParse(slugs["bank-id"]))
	if requiredBank == nil {
		json.NewEncoder(w).Encode(guru_errors.ReadBankFailed)
		panic(guru_errors.NewBankError(guru_errors.ReadBankFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(requiredBank)
	panic(guru_errors.NewBankError(guru_errors.ReadBankSuccess).GetSpecificMessage())
	// fmt.Println("Inside read bank by id controller function GET REQUEST DONE")
}

func UpdateBankById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside update bank by id controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var newBankTemp *bank_service.Bank
	err := json.NewDecoder(r.Body).Decode(&newBankTemp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewBankError(guru_errors.UpdateBankFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)
	updatedBank := adminObject.UpdateBankObject(uuid.MustParse(slugs["bank-id"]), newBankTemp)
	json.NewEncoder(w).Encode(updatedBank)
	panic(guru_errors.NewBankError(guru_errors.UpdateBankSuccess).GetSpecificMessage())
	// fmt.Println("Inside update bank by id controller function PUT REQUEST DONE")

}

func DeleteBankById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside delete bank by id controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	deletedBank := adminObject.DeleteBank(uuid.MustParse(slugs["bank-id"]))
	json.NewEncoder(w).Encode(deletedBank)
	panic(guru_errors.NewBankError(guru_errors.DeleteBankSuccess).GetSpecificMessage())
	// fmt.Println("Inside delete bank by id controller function DELETE REQUEST DONE")

}

func NetWorthEachBank(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside net worth each bank controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	netWorthOfEachBank := adminObject.GetNetWorthOfEachBank()
	if netWorthOfEachBank == nil {
		json.NewEncoder(w).Encode(guru_errors.NetWorthEachBankFailed)
		panic(guru_errors.NewBankError(guru_errors.NetWorthEachBankFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(netWorthOfEachBank)
	panic(guru_errors.NewBankError(guru_errors.NetWorthEachBankSuccess).GetSpecificMessage())
}
func NetWorthGivenBank(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside net worth given bank controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	slugs := mux.Vars(r)
	netWorthOfGivenBank := adminObject.GetNetWorthOfGivenBank(uuid.MustParse(slugs["bank-id"]))
	if netWorthOfGivenBank == nil {
		json.NewEncoder(w).Encode(guru_errors.NetWorthGivenBankFailed)
		panic(guru_errors.NewBankError(guru_errors.NetWorthGivenBankFailed).GetSpecificMessage())
	}
	json.NewEncoder(w).Encode(netWorthOfGivenBank)
	panic(guru_errors.NewBankError(guru_errors.NetWorthGivenBankSuccess).GetSpecificMessage())
}

func BankNameBalanceMapByBankId(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside bank name balance map by bank id controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewBankError(guru_errors.BankNameBalanceMapFailed).GetSpecificMessage())
	}
	slugs := mux.Vars(r)

	requiredBankBalanceMap := adminObject.BankTransferMapNameBalanceByBankId(
		uuid.MustParse(slugs["bank-id"]),
		newInputTransaction.StartDate,
		newInputTransaction.EndDate,
	)
	json.NewEncoder(w).Encode(requiredBankBalanceMap)
	panic(guru_errors.NewBankError(guru_errors.BankNameBalanceMapSuccess).GetSpecificMessage())

}

func BankNameBalanceMapAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Inside bank name balance map all controller function")
	var adminObject *customer_service.Customer = GetAdminObjectFromCookie(w, r)
	var newInputTransaction *InputTransaction
	err := json.NewDecoder(r.Body).Decode(&newInputTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(guru_errors.NewBankError(guru_errors.BankNameBalanceMapAllFailed).GetSpecificMessage())
	}
	requiredBankBalanceMap := adminObject.BankTransferMapNameBalanceAll(
		newInputTransaction.StartDate,
		newInputTransaction.EndDate,
	)
	json.NewEncoder(w).Encode(requiredBankBalanceMap)
	panic(guru_errors.NewBankError(guru_errors.BankNameBalanceMapAllSuccess).GetSpecificMessage())

}
