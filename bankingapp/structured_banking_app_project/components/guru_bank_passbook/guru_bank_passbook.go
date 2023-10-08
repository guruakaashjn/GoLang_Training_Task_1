package guru_bank_passbook

// import (
// 	"bankingapp/components/guru_bank_entry"
// 	"time"

// 	"github.com/google/uuid"
// )

// type BankPassbook struct {
// 	Entries []*guru_bank_entry.BankEntry
// }

// func NewBankPassbook() *BankPassbook {
// 	initialEntries := make([]*guru_bank_entry.BankEntry, 0)
// 	return &BankPassbook{
// 		Entries: initialEntries,
// 	}
// }

// func (bP *BankPassbook) ReadAllEntries(fromDate time.Time, toDate time.Time) (mapBankIdBalance map[uuid.UUID]int) {
// 	mapBankIdBalance = make(map[uuid.UUID]int)

// 	for i := 0; i < len(bP.Entries); i++ {
// 		bankEntryGoDate, _ := time.Parse("2006-01-02", bP.Entries[i].GetBankEntryTimeStampDate())
// 		if bankEntryGoDate.After(fromDate) && bankEntryGoDate.Before(toDate) {
// 			if bP.Entries[i].GetAmount() < 0 {
// 				mapBankIdBalance[bP.Entries[i].GetToBank()] += bP.Entries[i].GetAmount()
// 			} else {
// 				mapBankIdBalance[bP.Entries[i].GetFromBank()] += bP.Entries[i].GetAmount()
// 			}

// 		}
// 	}
// 	return mapBankIdBalance

// }

// func CreateBankPassbook() *BankPassbook {
// 	return NewBankPassbook()
// }

// func (bP *BankPassbook) AddEntry(fromBankId, toBankId uuid.UUID, amount int) {
// 	bP.Entries = append(bP.Entries, guru_bank_entry.CreateBankEntry(fromBankId, toBankId, amount))
// }
