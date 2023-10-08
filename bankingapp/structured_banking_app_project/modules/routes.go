package modules

import "bankingapp/app"

func RegisterModuleRoutes(app *app.App) {
	log := app.Log
	log.Print("Register Module routes.go")
	app.WG.Add(4)
	go registerCustomerRoutes(app)
	go registerBankRoutes(app)
	go registerAccountRoutes(app)
	go registerOfferRoutes(app)

	app.WG.Wait()
}
