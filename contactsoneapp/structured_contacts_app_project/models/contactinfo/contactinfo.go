package contactinfo

import (
	"contactsoneapp/models/contact"

	"github.com/jinzhu/gorm"
)

type ContactInfo struct {
	gorm.Model
	Contact      contact.Contact `gorm:"foreignkey:ContactRefer"`
	ContactRefer uint
	TypeName     string `json:"TypeName" gorm:"type:varchar(100)"`
	TypeValue    string `json:"TypeValue" gorm:"type:varchar(100)"`
	IsActive     bool   `json:"IsActive" gorm:"type:boolean;default:false"`
}
