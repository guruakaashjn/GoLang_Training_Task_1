package modules

import (
	"bankingapp/app"
	"bankingapp/components/guru_offer/controller"
	"bankingapp/components/guru_offer/service"
)

func registerOfferRoutes(appObj *app.App) {
	defer appObj.WG.Done()

	offerService := service.NewOfferService(appObj.DB, appObj.Repository)
	offerController := controller.NewOfferController(offerService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		offerController,
	})

}
