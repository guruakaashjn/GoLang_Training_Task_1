package guru_entries

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	TimeStampDate     string
	TimeStampTime     string
	SenderId          uuid.UUID
	ReceiverId        uuid.UUID
	SenderAccountId   uuid.UUID
	ReceiverAccountId uuid.UUID
	Amount            int
	TransactionType   string
}

func NewEntry(SenderId, ReceiverId uuid.UUID, SenderAccountId, ReceiverAccountId uuid.UUID, Amount int, TransactionType string) *Entry {
	now := time.Now()
	return &Entry{
		TimeStampDate:     now.Format("2006-01-02"),
		TimeStampTime:     now.Format("15:04:05"),
		SenderId:          SenderId,
		ReceiverId:        ReceiverId,
		SenderAccountId:   SenderAccountId,
		ReceiverAccountId: ReceiverAccountId,
		Amount:            Amount,
		TransactionType:   TransactionType,
	}
}

func CreateEntry(SenderId, ReceiverId uuid.UUID, SenderAccountId, ReceiverAccountId uuid.UUID, Amount int, TransactionType string) *Entry {
	return NewEntry(SenderId, ReceiverId, SenderAccountId, ReceiverAccountId, Amount, TransactionType)
}

func (e *Entry) GetTimeStampDate() string {
	return e.TimeStampDate
}
