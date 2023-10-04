package modules

import (
	"contactsoneapp/app"
	"contactsoneapp/components/guru_contacts/controller"
	"contactsoneapp/components/guru_contacts/service"
)

func registerContactRoutes(appObj *app.App) {
	defer appObj.WG.Done()
	contactService := service.NewContactService(appObj.DB, appObj.Repository)
	contactController := controller.NewContactController(contactService, appObj.Log)
	appObj.RegisterControllerRoutes([]app.Controller{
		contactController,
	})
}
