package guru_entries

import (
	"time"

	"github.com/google/uuid"
)

type Entries struct {
	timeStampDate     string
	timeStampTime     string
	senderId          uuid.UUID
	receiverId        uuid.UUID
	senderAccountId   uuid.UUID
	receiverAccountId uuid.UUID
	amount            int
	transactionType   string
}

func NewEntries(senderId, receiverId uuid.UUID, senderAccountId, receiverAccountId uuid.UUID, amount int, transactionType string) *Entries {
	now := time.Now()
	return &Entries{
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
