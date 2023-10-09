package controller

import (
	"bankingapp/components/guru_offer/service"
	"bankingapp/components/log"
	"bankingapp/middleware/auth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type OfferController struct {
	log     log.Log
	service *service.OfferService
}

func NewOfferController(offerService *service.OfferService, log log.Log) *OfferController {
	return &OfferController{
		service: offerService,
		log:     log,
	}
}

func (controller *OfferController) RegisterRoutes(router *mux.Router) {

	customerRouter := router.PathPrefix("/bank/{bank-id}/offer").Subrouter()
	customerRouter.HandleFunc("/", controller.RegisterOffer).Methods(http.MethodPost)
	customerRouter.HandleFunc("/", controller.GetAllOffers).Methods(http.MethodGet)
	customerRouter.HandleFunc("/{id}", controller.GetOfferById).Methods(http.MethodGet)
	customerRouter.Use(auth.IsAdmin)

	fmt.Println("[Offer register routes]")
}
