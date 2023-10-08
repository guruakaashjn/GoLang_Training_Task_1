package modules

import (
	"bankingapp/app"
	"bankingapp/components/guru_account/controller"
	"bankingapp/components/guru_account/service"
)

func registerAccountRoutes(appObj *app.App) {
	defer appObj.WG.Done()

	accountService := service.NewAccountService(appObj.DB, appObj.Repository)
	accountController := controller.NewAccountController(accountService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		accountController,
	})

}
