package modules

import (
	"bankingapp/app"
	"bankingapp/components/guru_bank/controller"
	"bankingapp/components/guru_bank/service"
)

func registerBankRoutes(appObj *app.App) {
	defer appObj.WG.Done()
	bankService := service.NewBankService(appObj.DB, appObj.Repository)
	bankController := controller.NewBankController(bankService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		bankController,
	})
}
