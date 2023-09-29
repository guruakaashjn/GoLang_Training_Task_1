package guru_bank_entry

import (
	"time"

	"github.com/google/uuid"
)

type BankEntry struct {
	timeStampDate string
	timeStampTime string
	fromBank      uuid.UUID
	toBank        uuid.UUID
	amount        int
}

func NewBankEntry(fromBankId, toBankId uuid.UUID, amount int) *BankEntry {
	now := time.Now()
	return &BankEntry{
		timeStampDate: now.Format("2006-01-02"),
		timeStampTime: now.Format("15:04:05"),
		fromBank:      fromBankId,
		toBank:        toBankId,
		amount:        amount,
	}
}
func (bE *BankEntry) GetBankEntryTimeStampDate() string {
	return bE.timeStampDate
}
func (bE *BankEntry) GetFromBank() uuid.UUID {
	return bE.fromBank
}

func (bE *BankEntry) GetToBank() uuid.UUID {
	return bE.toBank
}
func (bE *BankEntry) GetAmount() int {
	return bE.amount
}
func CreateBankEntry(fromBankId, toBankId uuid.UUID, amount int) *BankEntry {
	return NewBankEntry(fromBankId, toBankId, amount)
}
