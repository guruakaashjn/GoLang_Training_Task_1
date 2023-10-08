package controller

import (
	"bankingapp/errors"
	"bankingapp/models/offer"
	"bankingapp/web"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (controller *OfferController) RegisterOffer(w http.ResponseWriter, r *http.Request) {
	newOffer := offer.Offer{}
	err := web.UnmarshalJSON(r, &newOffer)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	slugs := mux.Vars(r)
	idTemp, _ := strconv.Atoi(slugs["bank-id"])

	newOffer.BankID = uint(idTemp)
	err = controller.service.RegisterOffer(&newOffer)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
	}
	web.RespondJSON(w, http.StatusCreated, newOffer)
}

func (controller *OfferController) GetAllOffers(w http.ResponseWriter, r *http.Request) {
	slugs := mux.Vars(r)
	idTemp, err := strconv.Atoi(slugs["bank-id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	limit, offset := web.ParseLimitAndOffset(r)
	givenAssociations := web.ParsePreloading(r)

	allOffers := &[]offer.Offer{}
	var totalCount int
	err = controller.service.GetAllOffers(allOffers, uint(idTemp), &totalCount, limit, offset, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allOffers)
}
func (controller *OfferController) GetOfferById(w http.ResponseWriter, r *http.Request) {
	requiredOffer := offer.Offer{}
	slugs := mux.Vars(r)
	bankIdTemp, err := strconv.Atoi(slugs["bank-id"])
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
	requiredOffer.ID = uint(idTemp)
	requiredOffer.BankID = uint(bankIdTemp)

	err = controller.service.GetOfferById(&requiredOffer, givenAssociations)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	web.RespondJSONWithXTotalCount(w, http.StatusOK, 1, requiredOffer)
}
