package modules

import (
	"contactsoneapp/app"
	"contactsoneapp/components/guru_user/controller"
	"contactsoneapp/components/guru_user/service"
)

func registerUserRoutes(appObj *app.App) {
	defer appObj.WG.Done()
	userService := service.NewUserService(appObj.DB, appObj.Repository)
	userController := controller.NewUserController(userService, appObj.Log)

	appObj.RegisterControllerRoutes([]app.Controller{
		userController,
	})
}
