package contactinfo

import (
	"contactsoneapp/components/log"
	"sync"

	"github.com/jinzhu/gorm"
)

type ModuleConfig struct {
	DB *gorm.DB
}

func NewContactInfoModuleConfig(db *gorm.DB) *ModuleConfig {
	return &ModuleConfig{
		DB: db,
	}
}

func (config *ModuleConfig) TableMigration(wg *sync.WaitGroup) {
	var models []interface{} = []interface{}{
		&ContactInfo{},
	}

	for _, model := range models {
		if err := config.DB.AutoMigrate(model).Error; err != nil {
			log.GetLogger().Print("Auto Migration ==> %s", err)
		}
	}
	if err := config.DB.Model(&ContactInfo{}).
		AddForeignKey("contact_refer", "contacts(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.GetLogger().Print("Auto Migration ==> %s", err)
	}

	log.GetLogger().Print("Contact Info Module Configured.")
}
