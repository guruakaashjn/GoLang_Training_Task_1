package modules

import (
	"bankingapp/app"
	"bankingapp/models/account"
	"bankingapp/models/bank"
	"bankingapp/models/bank_entry"
	"bankingapp/models/bank_passbook"
	"bankingapp/models/customer"
	"bankingapp/models/entry"
	"bankingapp/models/offer"
	"bankingapp/models/passbook"
)

func Configure(appObj *app.App) {
	customerModule := customer.NewCustomerModuleConfig(appObj.DB)

	bankModule := bank.NewBankModuleConfig(appObj.DB)
	bankPassbookModule := bank_passbook.NewBankPassbookModuleConfig(appObj.DB)
	bankEntryModule := bank_entry.NewBankEntryModuleConfig(appObj.DB)

	accountModule := account.NewAccountModuleConfig(appObj.DB)
	passbookModule := passbook.NewPassbookModuleConfig(appObj.DB)
	entryModule := entry.NewEntryModuleConfig(appObj.DB)

	offerModule := offer.NewOfferModuleConfig(appObj.DB)

	appObj.MigrateTables([]app.ModuleConfig{
		customerModule,
		bankModule,
		bankPassbookModule,
		bankEntryModule,
		accountModule,
		passbookModule,
		entryModule,
		offerModule,
	})

}
