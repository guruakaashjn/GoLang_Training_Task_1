package guru_entries

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	timeStampDate     string
	timeStampTime     string
	senderId          uuid.UUID
	receiverId        uuid.UUID
	senderAccountId   uuid.UUID
	receiverAccountId uuid.UUID
	amount            int
	transactionType   string
}

func NewEntry(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) *Entry {
	now := time.Now()
	return &Entry{
		timeStampDate:     now.Format("2006-01-02"),
		timeStampTime:     now.Format("15:04:05"),
		senderId:          senderId,
		receiverId:        receiverId,
		senderAccountId:   senderAccountId,
		receiverAccountId: receiverAccountId,
		amount:            amount,
		transactionType:   transactionType,
	}
}

func CreateEntry(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) *Entry {
	return NewEntry(senderId, receiverId, senderAccountId, receiverAccountId, amount, transactionType)
}

func (e *Entry) GetTimeStampDate() string {
	return e.timeStampDate
}
