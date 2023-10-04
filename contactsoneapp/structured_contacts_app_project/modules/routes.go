package modules

import "contactsoneapp/app"

func RegisterModuleRoutes(app *app.App) {
	log := app.Log
	log.Print("Register Module routes.go")
	app.WG.Add(3)
	go registerUserRoutes(app)
	go registerContactRoutes(app)
	go registerContactInfoRoutes(app)

	app.WG.Wait()
}
