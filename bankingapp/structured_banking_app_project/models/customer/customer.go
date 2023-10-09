package customer

import (
	"bankingapp/models/account"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string `json:"FirstName" gorm:"type:varchar(100)"`
	LastName  string `json:"LastName" gorm:"type:varchar(100)"`
	UserName  string `json:"UserName" gorm:"type:varchar(100)"`
	Password  string `json:"Password" gorm:"type:varchar(100)"`
	// TotalBalance int    `json:"TotalBalance" gorm:"type:int;default:0"`
	IsAdmin  bool `json:"IsAdmin" gorm:"type:bool;default:false"`
	IsActive bool `json:"IsActive" gorm:"type:bool;default:true"`
	Accounts []account.Account
}

type CustomerDTO struct {
	FirstName string `json:"FirstName" gorm:"type:varchar(100)"`
	LastName  string `json:"LastName" gorm:"type:varchar(100)"`
	UserName  string `json:"UserName" gorm:"type:varchar(100)"`
	// TotalBalance int    `json:"TotalBalance" gorm:"type:int;default:0"`
	IsActive bool `json:"IsActive" gorm:"type:bool;default:true"`
	Accounts []account.Account
}
