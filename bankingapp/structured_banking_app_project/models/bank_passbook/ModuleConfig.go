package bank_passbook

import (
	"bankingapp/components/log"

	"github.com/jinzhu/gorm"
)

type ModuleConfig struct {
	DB *gorm.DB
}

func NewBankPassbookModuleConfig(db *gorm.DB) *ModuleConfig {
	return &ModuleConfig{
		DB: db,
	}
}

func (config *ModuleConfig) TableMigration() {
	var models []interface{} = []interface{}{
		&BankPassbook{},
	}

	for _, model := range models {
		if err := config.DB.AutoMigrate(model).Error; err != nil {
			log.GetLogger().Print("Auto Migration ==> %s", err)
		}
	}
	log.GetLogger().Print("Bank Passbook module configured....")
}
