package guru_bank

import (
	account_service "bankingapp/components/guru_account/service"
	"bankingapp/components/guru_bank_passbook"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var Banks = make([]*Bank, 0)

type Bank struct {
	BankId       uuid.UUID
	FullName     string
	Abbreviation string
	IsActive     bool
	Accounts     []*account_service.Account
	bankPassbook *guru_bank_passbook.BankPassbook
}

func NewBank(FullName string) *Bank {
	var abbr string = setAbbreviation(FullName)
	var initialAccountsList []*account_service.Account = make([]*account_service.Account, 0)

	var bankPassbookInitial *guru_bank_passbook.BankPassbook = guru_bank_passbook.CreateBankPassbook()
	var newBankObject = &Bank{
		BankId:       uuid.New(),
		FullName:     FullName,
		Abbreviation: abbr,
		IsActive:     true,
		Accounts:     initialAccountsList,

		bankPassbook: bankPassbookInitial,
	}
	Banks = append(Banks, newBankObject)
	return newBankObject
}

func (b *Bank) GetAbbreviation() string {
	return b.Abbreviation
}

func setAbbreviation(FullName string) (abbr string) {
	abbr = FullName[0:4] + FullName[len(FullName)-4:]
	max := rand.Intn(10000-5000) + 5000
	min := rand.Intn(4999-0) + 0
	rnd := rand.Intn(max-min) + min
	abbr += strconv.Itoa(rnd)

	if checkAbbreviation(abbr) {
		return abbr
	}

	return setAbbreviation(FullName)
}

func checkAbbreviation(abbr string) (flag bool) {
	for i := 0; i < len(Banks); i++ {
		if Banks[i].GetAbbreviation() == abbr {
			return false
		}

	}
	return true
}

func CreateBank(FullName string) (bank *Bank) {

	return NewBank(FullName)

}

func (b *Bank) ReadBank() (bool, *Bank) {
	if b.IsActive {
		return true, b
	}
	return false, b

}
func ReadBankById(bankIdTemp uuid.UUID) (bool, *Bank) {
	var bank *Bank
	for i := 0; i < len(Banks); i++ {
		if Banks[i].BankId == bankIdTemp {
			bank = Banks[i]
			break
		}
	}
	if bank.IsActive {
		return true, bank
	}
	return false, bank
}
func ReadAllBanks() []*Bank {

	var allBanks []*Bank
	for i := 0; i < len(Banks); i++ {
		if Banks[i].IsActive {
			allBanks = append(allBanks, Banks[i])
		}
	}

	return allBanks

}

func (b *Bank) UpdateBank(updateValue string) *Bank {

	b.FullName = updateValue
	b.Abbreviation = setAbbreviation(b.FullName)
	return b

}

func (b *Bank) UpdateBankObject(bankTempObject *Bank) *Bank {
	if bankTempObject.FullName != "" && bankTempObject.FullName != b.FullName {
		b.FullName = bankTempObject.FullName
		b.Abbreviation = setAbbreviation(b.FullName)
	}

	return b

}

func (b *Bank) DeleteBank() *Bank {

	b.IsActive = false
	return b

}

// HELPER FUNCTIONS
func (b *Bank) GetBankId() uuid.UUID {
	return b.BankId
}
func (b *Bank) GetBankName() string {
	return b.FullName
}
func (b *Bank) GetNetWorthOfBank() (networth int) {

	for i := 0; i < len(b.Accounts); i++ {
		networth += b.Accounts[i].GetBalance()
	}

	return networth
}

func (b *Bank) CheckBankContainsActiveAccounts() bool {
	var flag bool = false
	for i := 0; i < len(b.Accounts); i++ {
		if b.Accounts[i].GetIsActive() {
			flag = true
			break
		}
	}
	return flag
}

func (b *Bank) GetBankPassbook() *guru_bank_passbook.BankPassbook {
	return b.bankPassbook
}

func (b *Bank) ReadPassbookFromRange(fromDate time.Time, toDate time.Time) map[uuid.UUID]int {

	return b.bankPassbook.ReadAllEntries(fromDate, toDate)
}
