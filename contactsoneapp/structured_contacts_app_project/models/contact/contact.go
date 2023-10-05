package contact

import (
	"contactsoneapp/models/contactinfo"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	UserID       uint
	FirstName    string                    `json:"FirstName" gorm:"type:varchar(100)"`
	LastName     string                    `json:"LastName" gorm:"type:varchar(100)"`
	IsActive     bool                      `json:"IsActive" gorm:"type:boolean;default:false"`
	ContactInfos []contactinfo.ContactInfo `gorm:"foreignkey:ContactRefer"`
}
