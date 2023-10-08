package bank_passbook

import (
	"bankingapp/models/bank_entry"

	"github.com/jinzhu/gorm"
)

type BankPassbook struct {
	gorm.Model
	BankID  uint
	Entries []bank_entry.BankEntry
}

func NewBankPassbook() *BankPassbook {
	initialEntries := make([]bank_entry.BankEntry, 0)
	return &BankPassbook{
		Entries: initialEntries,
	}
}
