package guru_passbook

// import (
// 	"bankingapp/components/guru_entries"
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// )

// type Passbook struct {
// 	Entries []*guru_entries.Entry
// }

// func NewPassbook(senderId uint, senderAccountId uint, balance int) *Passbook {

// 	var newEntries []*guru_entries.Entry = make([]*guru_entries.Entry, 0)
// 	newEntries = append(newEntries, guru_entries.CreateEntry(senderId, senderId, senderAccountId, senderAccountId, balance, "CREDIT"))
// 	fmt.Println("New Entries in NewPassbook in guru_passbook.go", newEntries)
// 	return &Passbook{
// 		Entries: newEntries,
// 	}

// }

// func (p *Passbook) AddEntry(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) {
// 	var newEntry *guru_entries.Entry = guru_entries.CreateEntry(senderId, receiverId, senderAccountId, receiverAccountId, amount, transactionType)
// 	p.Entries = append(p.Entries, newEntry)

// }

// func (p *Passbook) ReadPassbook(startDate, endDate time.Time) (requiredPassbookRange *Passbook) {

// 	var requiredPassbookRangeEntries []*guru_entries.Entry
// 	for i := 0; i < len(p.Entries); i++ {
// 		entryDate, _ := time.Parse("2006-01-02", p.Entries[i].GetTimeStampDate())

// 		if entryDate.After(startDate) && entryDate.Before(endDate) {
// 			requiredPassbookRangeEntries = append(requiredPassbookRangeEntries, p.Entries[i])
// 		}
// 	}
// 	requiredPassbookRange = &Passbook{
// 		Entries: requiredPassbookRangeEntries,
// 	}
// 	return requiredPassbookRange

// }
