package contactinfo

import (
	"github.com/jinzhu/gorm"
)

type ContactInfo struct {
	gorm.Model

	TypeName     string `json:"TypeName" gorm:"type:varchar(100)"`
	TypeValue    string `json:"TypeValue" gorm:"type:varchar(100)"`
	IsActive     bool   `json:"IsActive" gorm:"type:boolean;default:false"`
	ContactRefer uint
}
