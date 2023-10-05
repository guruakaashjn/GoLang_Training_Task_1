package user

import (
	"contactsoneapp/models/contact"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"FirstName" gorm:"type:varchar(100)"`
	LastName  string `json:"LastName" gorm:"type:varchar(100)"`
	UserName  string `json:"UserName" gorm:"type:varchar(500)"`
	Password  string `json:"Password" gorm:"type:varchar(500)"`
	IsAdmin   bool   `json:"IsAdmin" gorm:"type:boolean;default:false"`
	IsActive  bool   `json:"IsActive" gorm:"type:boolean;default:true"`
	Contacts  []contact.Contact
}
