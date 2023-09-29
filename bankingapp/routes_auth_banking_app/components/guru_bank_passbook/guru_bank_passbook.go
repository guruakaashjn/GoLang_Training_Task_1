package guru_bank_passbook

import (
	"bankingapp/components/guru_bank_entry"
	"time"

	"github.com/google/uuid"
)

type BankPassbook struct {
	entries []*guru_bank_entry.BankEntry
}

func NewBankPassbook() *BankPassbook {
	initialEntries := make([]*guru_bank_entry.BankEntry, 0)
	return &BankPassbook{
		entries: initialEntries,
	}
}

func (bP *BankPassbook) ReadAllEntries(fromDate time.Time, toDate time.Time) (mapBankIdBalance map[uuid.UUID]int) {
	mapBankIdBalance = make(map[uuid.UUID]int)

	for i := 0; i < len(bP.entries); i++ {
		bankEntryGoDate, _ := time.Parse("2006-01-02", bP.entries[i].GetBankEntryTimeStampDate())
		if bankEntryGoDate.After(fromDate) && bankEntryGoDate.Before(toDate) {

			mapBankIdBalance[bP.entries[i].GetToBank()] += bP.entries[i].GetAmount()
		}
	}
	return mapBankIdBalance

}

func CreateBankPassbook() *BankPassbook {
	return NewBankPassbook()
}

func (bP *BankPassbook) AddEntry(fromBankId, toBankId uuid.UUID, amount int) {
	bP.entries = append(bP.entries, guru_bank_entry.CreateBankEntry(fromBankId, toBankId, amount))
}
