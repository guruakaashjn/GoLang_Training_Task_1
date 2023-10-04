package modules

import (
	"contactsoneapp/app"
	"contactsoneapp/components/guru_contact_details/controller"
	"contactsoneapp/components/guru_contact_details/service"
)

func registerContactInfoRoutes(appObj *app.App) {
	defer appObj.WG.Done()
	contactInfoService := service.NewContactInfoService(appObj.DB, appObj.Repository)
	contactInfoController := controller.NewContactInfoController(contactInfoService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		contactInfoController,
	})
}
