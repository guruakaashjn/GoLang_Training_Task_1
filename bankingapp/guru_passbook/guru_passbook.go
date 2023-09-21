package guru_passbook

import (
	"bankingapp/guru_entries"
)

type Passbook struct {
	entries []*guru_entries.Entries
}

func NewPassbook() *Passbook {
	return &Passbook{
		entries: make([]*guru_entries.Entries, 0),
	}

}
