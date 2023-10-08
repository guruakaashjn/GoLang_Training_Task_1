package guru_entries

// import (
// 	"time"
// )

// type Entry struct {
// 	TimeStampDate     string
// 	TimeStampTime     string
// 	SenderId          uint
// 	ReceiverId        uint
// 	SenderAccountId   uint
// 	ReceiverAccountId uint
// 	Amount            int
// 	TransactionType   string
// 	PassbookID        uint
// }

// func NewEntry(SenderId, ReceiverId uint, SenderAccountId, ReceiverAccountId uint, Amount int, TransactionType string) *Entry {
// 	now := time.Now()
// 	return &Entry{
// 		TimeStampDate:     now.Format("2006-01-02"),
// 		TimeStampTime:     now.Format("15:04:05"),
// 		SenderId:          SenderId,
// 		ReceiverId:        ReceiverId,
// 		SenderAccountId:   SenderAccountId,
// 		ReceiverAccountId: ReceiverAccountId,
// 		Amount:            Amount,
// 		TransactionType:   TransactionType,
// 	}
// }

// func CreateEntry(SenderId, ReceiverId uint, SenderAccountId, ReceiverAccountId uint, Amount int, TransactionType string) *Entry {
// 	return NewEntry(SenderId, ReceiverId, SenderAccountId, ReceiverAccountId, Amount, TransactionType)
// }

// func (e *Entry) GetTimeStampDate() string {
// 	return e.TimeStampDate
// }
