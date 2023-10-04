package contact

import (
	"contactsoneapp/models/user"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	User      user.User
	UserID    uint
	FirstName string `json:"FirstName" gorm:"type:varchar(100)"`
	LastName  string `json:"LastName" gorm:"type:varchar(100)"`
	IsActive  bool   `json:"IsActive" gorm:"type:boolean;default:false"`
}
