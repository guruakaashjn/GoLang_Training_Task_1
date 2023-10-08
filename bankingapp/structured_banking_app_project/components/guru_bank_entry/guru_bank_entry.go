package guru_bank_entry

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type BankEntry struct {
// 	TimeStampDate string
// 	TimeStampTime string
// 	FromBank      uuid.UUID
// 	ToBank        uuid.UUID
// 	Amount        int
// }

// func NewBankEntry(fromBankId, toBankId uuid.UUID, Amount int) *BankEntry {
// 	now := time.Now()
// 	return &BankEntry{
// 		TimeStampDate: now.Format("2006-01-02"),
// 		TimeStampTime: now.Format("15:04:05"),
// 		FromBank:      fromBankId,
// 		ToBank:        toBankId,
// 		Amount:        Amount,
// 	}
// }
// func (bE *BankEntry) GetBankEntryTimeStampDate() string {
// 	return bE.TimeStampDate
// }
// func (bE *BankEntry) GetFromBank() uuid.UUID {
// 	return bE.FromBank
// }

// func (bE *BankEntry) GetToBank() uuid.UUID {
// 	return bE.ToBank
// }
// func (bE *BankEntry) GetAmount() int {
// 	return bE.Amount
// }
// func CreateBankEntry(fromBankId, toBankId uuid.UUID, Amount int) *BankEntry {
// 	return NewBankEntry(fromBankId, toBankId, Amount)
// }
