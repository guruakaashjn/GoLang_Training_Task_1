package guru_passbook

import (
	"bankingapp/guru_entries"
	"time"

	"github.com/google/uuid"
)

type Passbook struct {
	entries []*guru_entries.Entry
}

func NewPassbook(senderId uuid.UUID, senderAccountId uuid.UUID, balance int) *Passbook {

	var newEntries []*guru_entries.Entry = make([]*guru_entries.Entry, 0)
	newEntries = append(newEntries, guru_entries.CreateEntry(senderId, senderId, senderAccountId, senderAccountId, balance, "CREDIT"))
	return &Passbook{
		entries: newEntries,
	}

}

func (p *Passbook) AddEntry(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) {
	var newEntry *guru_entries.Entry = guru_entries.CreateEntry(senderId, receiverId, senderAccountId, receiverAccountId, amount, transactionType)
	p.entries = append(p.entries, newEntry)

}

func (p *Passbook) ReadPassbook(startDate, endDate time.Time) (requiredPassbookRange *Passbook) {

	var requiredPassbookRangeEntries []*guru_entries.Entry
	for i := 0; i < len(p.entries); i++ {
		entryDate, _ := time.Parse("2006-01-02", p.entries[i].GetTimeStampDate())

		if entryDate.After(startDate) && entryDate.Before(endDate) {
			requiredPassbookRangeEntries = append(requiredPassbookRangeEntries, p.entries[i])
		}
	}
	requiredPassbookRange = &Passbook{
		entries: requiredPassbookRangeEntries,
	}
	return requiredPassbookRange

}
