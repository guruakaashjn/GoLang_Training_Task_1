package passbook

import (
	"bankingapp/models/entry"

	"github.com/jinzhu/gorm"
)

type Passbook struct {
	gorm.Model
	AccountID uint
	Entries   []entry.Entry
}

func NewPassbook(senderId uint, senderAccountId uint, balance int) *Passbook {
	var newEntries []entry.Entry = make([]entry.Entry, 0)

	return &Passbook{
		Entries: newEntries,
	}
}
