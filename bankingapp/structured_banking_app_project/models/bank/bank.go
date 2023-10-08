package bank

import (
	"bankingapp/models/account"
	"bankingapp/models/bank_passbook"
	"bankingapp/models/offer"

	"github.com/jinzhu/gorm"
)

type Bank struct {
	gorm.Model
	FullName     string `json:"FullName" gorm:"type:varchar(100)"`
	Abbreviation string `json:"Abbreviation" gorm:"type:varchar(100)"`
	IsActive     bool   `json:"IsActive" gorm:"type:bool;default:true"`
	Accounts     []account.Account
	BankPassbook bank_passbook.BankPassbook
	Offers       []offer.Offer
}
