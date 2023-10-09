package controller

import (
	"bankingapp/errors"
	"bankingapp/models/bank"
	"bankingapp/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (controller *BankController) RegisterBank(w http.ResponseWriter, r *http.Request) {
	newBank := bank.Bank{}
	err := web.UnmarshalJSON(r, &newBank)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	err = controller.service.CreateBank(&newBank)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newBank)
}

func (controller *BankController) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	allBanks := &[]bank.Bank{}

	var totalCount int
	limit, offset, err := web.ParseLimitAndOffset(r)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	givenAssociations := web.ParsePreloading(r)
	err = controller.service.GetAllBanks(allBanks, &totalCount, limit, offset, givenAssociations)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allBanks)
}
func (controller *BankController) GetBankById(w http.ResponseWriter, r *http.Request) {
	requiredBank := bank.Bank{}

	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	givenAssociations := web.ParsePreloading(r)
	err = controller.service.GetBankById(&requiredBank, idTemp, givenAssociations)

	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredBank)
}
func (controller *BankController) UpdateBank(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bank to update")
	bankToUpdate := bank.Bank{}

	fmt.Println(r.Body)

	err := web.UnmarshalJSON(r, &bankToUpdate)
	if err != nil {
		fmt.Println("error from unmarshal JSON")
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)

	intId, err := strconv.Atoi(vars["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bankToUpdate.ID = uint(intId)

	fmt.Println("Bank to update")
	fmt.Println(&bankToUpdate)
	err = controller.service.UpdateBank(&bankToUpdate)

	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}

	web.RespondJSON(w, http.StatusOK, bankToUpdate)

}
func (controller *BankController) DeleteBank(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------------------Delete Bank controller")
	controller.log.Print("Delete bank call")
	bankToDelete := bank.Bank{}

	var err error
	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bankToDelete.ID = uint(intId)
	// fmt.Println("----------------------------1")
	err = controller.service.DeleteBank(&bankToDelete)
	// fmt.Println("----------------------------2")
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	// fmt.Println("----------------------------3")
	web.RespondJSON(w, http.StatusOK, "Delete  successful.")
}

func (controller *BankController) AllBankNetWorth(w http.ResponseWriter, r *http.Request) {
	mapAllBankNetWorth := make(map[string]int, 0)

	err := controller.service.AllBankNetWorth(mapAllBankNetWorth)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, mapAllBankNetWorth)
}

func (controller *BankController) BankNetworth(w http.ResponseWriter, r *http.Request) {

	var err error
	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	requiredBank := &bank.Bank{}
	requiredBank.ID = uint(intId)

	mapBankNetWorth := make(map[string]int, 0)

	err = controller.service.BankNetWorth(requiredBank, mapBankNetWorth)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, mapBankNetWorth)
}

func (controller *BankController) BankBalanceMap(w http.ResponseWriter, r *http.Request) {
	var err error

	startDateTimeDotTime, endDateTimeDotTime, err := web.ParseStartDateEndDate(r)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	slugs := mux.Vars(r)
	intId, err := strconv.Atoi(slugs["id"])
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	requiredBank := &bank.Bank{}
	requiredBank.ID = uint(intId)

	mapBankBalance := make(map[string]map[string]int, 0)

	err = controller.service.BankBalanceMap(requiredBank, mapBankBalance, startDateTimeDotTime, endDateTimeDotTime)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, mapBankBalance)

}

func (controller *BankController) AllBankBalanceMap(w http.ResponseWriter, r *http.Request) {

	startDateTimeDotTime, endDateTimeDotTime, err := web.ParseStartDateEndDate(r)
	if err != nil {
		// controller.log.Print(err)
		controller.log.PrintError(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	mapBankBalance := make([]map[string]map[string]int, 0)
	err = controller.service.AllBankBalanceMap(&mapBankBalance, startDateTimeDotTime, endDateTimeDotTime)
	if err != nil {
		// controller.log.Print(err.Error())
		controller.log.PrintError(err)
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, mapBankBalance)

}
