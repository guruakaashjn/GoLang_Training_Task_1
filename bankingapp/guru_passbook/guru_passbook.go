package guru_passbook

import (
	"bankingapp/guru_entries"

	"github.com/google/uuid"
)

type Passbook struct {
	entries []*guru_entries.Entries
}

func NewPassbook(senderId uuid.UUID, senderAccountId uuid.UUID, balance int) *Passbook {

	var newEntries []*guru_entries.Entries = make([]*guru_entries.Entries, 0)
	newEntries = append(newEntries, guru_entries.NewEntries(senderId, senderId, senderAccountId, senderAccountId, balance, "CREDIT"))
	return &Passbook{
		entries: newEntries,
	}

}

func (p *Passbook) AddEntry(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) {
	var newEntry *guru_entries.Entries = guru_entries.NewEntries(senderId, receiverId, senderAccountId, receiverAccountId, amount, transactionType)
	p.entries = append(p.entries, newEntry)

}
