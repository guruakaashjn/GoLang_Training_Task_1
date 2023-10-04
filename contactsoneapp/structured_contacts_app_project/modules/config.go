package modules

import (
	"contactsoneapp/app"
	"contactsoneapp/models/contact"
	"contactsoneapp/models/contactinfo"
	"contactsoneapp/models/user"
)

func Configure(appObj *app.App) {
	userModule := user.NewUserModuleConfig(appObj.DB)
	contactModule := contact.NewContactModuleConfig(appObj.DB)
	contactInfoModule := contactinfo.NewContactInfoModuleConfig(appObj.DB)

	appObj.MigrateTables([]app.ModuleConfig{userModule, contactModule, contactInfoModule})

}
