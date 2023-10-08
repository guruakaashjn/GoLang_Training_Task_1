package bank_entry

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BankEntry struct {
	gorm.Model
	TimeStampDate   string
	TimeStampTime   string
	FromBank        uint
	ToBank          uint
	Amount          int
	TransactionType string
	BankPassbookID  uint
}

func NewBankEntry(fromBankId, toBankId uint, Amount int, transactionType string) *BankEntry {
	now := time.Now()
	return &BankEntry{
		TimeStampDate:   now.Format("2006-01-02"),
		TimeStampTime:   now.Format("15:04:05"),
		FromBank:        fromBankId,
		ToBank:          toBankId,
		Amount:          Amount,
		TransactionType: transactionType,
	}
}
