package modules

import (
	"bankingapp/app"
	"bankingapp/components/guru_customer/controller"
	"bankingapp/components/guru_customer/service"
)

func registerCustomerRoutes(appObj *app.App) {
	defer appObj.WG.Done()

	customerService := service.NewCustomerService(appObj.DB, appObj.Repository)
	customerController := controller.NewCustomerController(customerService, appObj.Log)
	appObj.RegisterControllerRoutes([]app.Controller{
		customerController,
	})

}
