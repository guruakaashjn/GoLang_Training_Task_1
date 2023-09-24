package guru_bank

import (
	"bankingapp/guru_account"
	"bankingapp/guru_bank_passbook"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var Banks = make([]*Bank, 0)

type Bank struct {
	bankId       uuid.UUID
	fullName     string
	abbreviation string
	isActive     bool
	Accounts     []*guru_account.Account
	bankPassbook *guru_bank_passbook.BankPassbook
}

func NewBank(fullName string) *Bank {
	var abbr string = setAbbreviation(fullName)
	var initialAccountsList []*guru_account.Account = make([]*guru_account.Account, 0)

	var bankPassbookInitial *guru_bank_passbook.BankPassbook = guru_bank_passbook.CreateBankPassbook()
	var newBankObject = &Bank{
		bankId:       uuid.New(),
		fullName:     fullName,
		abbreviation: abbr,
		isActive:     true,
		Accounts:     initialAccountsList,

		bankPassbook: bankPassbookInitial,
	}
	Banks = append(Banks, newBankObject)
	return newBankObject
}

func (b *Bank) GetAbbreviation() string {
	return b.abbreviation
}

func setAbbreviation(fullName string) (abbr string) {
	abbr = fullName[0:4] + fullName[len(fullName)-4:]
	max := rand.Intn(10000-5000) + 5000
	min := rand.Intn(4999-0) + 0
	rnd := rand.Intn(max-min) + min
	abbr += strconv.Itoa(rnd)

	if checkAbbreviation(abbr) {
		return abbr
	}

	return setAbbreviation(fullName)
}

func checkAbbreviation(abbr string) (flag bool) {
	for i := 0; i < len(Banks); i++ {
		if Banks[i].GetAbbreviation() == abbr {
			return false
		}

	}
	return true
}

func CreateBank(fullName string) (bank *Bank) {

	return NewBank(fullName)

}

func (b *Bank) ReadBank() (bool, *Bank) {
	if b.isActive {
		return true, b
	}
	return false, b

}
func ReadBankById(bankIdTemp uuid.UUID) (bool, *Bank) {
	var bank *Bank
	for i := 0; i < len(Banks); i++ {
		if Banks[i].bankId == bankIdTemp {
			bank = Banks[i]
			break
		}
	}
	if bank.isActive {
		return true, bank
	}
	return false, bank
}
func ReadAllBanks() []*Bank {

	var allBanks []*Bank
	for i := 0; i < len(Banks); i++ {
		if Banks[i].isActive {
			allBanks = append(allBanks, Banks[i])
		}
	}

	return allBanks

}

func (b *Bank) UpdateBank(updateValue string) *Bank {

	b.fullName = updateValue
	b.abbreviation = setAbbreviation(b.fullName)
	return b

}

func (b *Bank) DeleteBank() *Bank {

	b.isActive = false
	return b

}

// HELPER FUNCTIONS
func (b *Bank) GetBankId() uuid.UUID {
	return b.bankId
}
func (b *Bank) GetBankName() string {
	return b.fullName
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
