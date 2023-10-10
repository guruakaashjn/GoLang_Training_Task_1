package account

import (
	"bankingapp/models/offer"
	"bankingapp/models/passbook"

	"github.com/jinzhu/gorm"
)

type Rough interface {
	RoughFunc()
}
type Account struct {
	gorm.Model
	IsActive   bool `json:"Bank" gorm:"type:boolean;default:true"`
	Balance    int  `json:"Balance" gorm:"type:int"`
	BankID     uint
	CustomerID uint
	Passbook   passbook.Passbook
	Offers     []offer.Offer `gorm:"many2many:account_offers_join;"`
}

type AccountDTO struct {
	IsActive   bool `json:"Bank" gorm:"type:boolean;default:true"`
	Balance    int  `json:"Balance" gorm:"type:int"`
	BankID     uint
	CustomerID uint
	Passbook   passbook.Passbook
	Offers     []offer.Offer `gorm:"many2many:account_offers_join;"`
}

func (a *Account) RoughFunc() {

}
func (a *AccountDTO) RoughFunc() {

}
