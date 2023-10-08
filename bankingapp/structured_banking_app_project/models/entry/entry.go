package entry

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Entry struct {
	gorm.Model
	TimeStampDate     string
	TimeStampTime     string
	SenderId          uint
	ReceiverId        uint
	SenderAccountId   uint
	ReceiverAccountId uint
	Amount            int
	TransactionType   string
	PassbookID        uint
}

func NewEntry(SenderId, ReceiverId, SenderAccountId, ReceiverAccountId uint, Amount int, TransactionType string) *Entry {
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
