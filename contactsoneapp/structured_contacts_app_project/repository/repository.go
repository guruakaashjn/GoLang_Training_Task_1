package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll(uow *UnitOfWork, out interface{}, queryProcessor ...QueryProcessor) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	UpdateWithMap(upw *UnitOfWork, model interface{}, value map[string]interface{}, queryProcessors ...QueryProcessor) error
	GetRecordForId(uow *UnitOfWork, id uint, out interface{}, queryProcessor ...QueryProcessor) error
	Delete(uow *UnitOfWork, out interface{}) error
	Save(uow *UnitOfWork, value interface{}) error
	GetRecord(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
}

type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

type UnitOfWork struct {
	DB       *gorm.DB
	Commited bool
	Readonly bool
}

func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	commit := false
	if readonly {
		return &UnitOfWork{
			DB:       db.New(),
			Commited: commit,
			Readonly: readonly,
		}
	}
	return &UnitOfWork{
		DB:       db.New().Begin(),
		Commited: commit,
		Readonly: readonly,
	}
}

func (uow *UnitOfWork) Commit() {
	if !uow.Readonly && !uow.Commited {
		uow.Commited = true
		uow.DB.Commit()
	}
}
func (uow *UnitOfWork) RollBack() {
	if !uow.Commited && !uow.Readonly {
		uow.DB.Rollback()
	}
}

func (repository *GormRepository) Save(uow *UnitOfWork, value interface{}) error {
	return uow.DB.Debug().Save(value).Error
}

func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Where(condition, args...)
		return db, nil
	}
}

func Select(query interface{}, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Select(query, args...)
		return db, nil
	}
}

func executeQueryProcessors(db *gorm.DB, out interface{}, queryProcessors ...QueryProcessor) (*gorm.DB, error) {
	var err error
	for _, query := range queryProcessors {
		if query != nil {
			db, err = query(db, out)
			if err != nil {
				return db, err
			}
		}
	}
	return db, nil
}

func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().Find(out).Error
}
func (repository *GormRepository) GetRecordForId(uow *UnitOfWork, id uint, out interface{}, queryProcessors ...QueryProcessor) error {
	queryProcessors = append([]QueryProcessor{Filter("id = ?", id)}, queryProcessors...)
	return repository.GetRecord(uow, out, queryProcessors...)
}

func (repository *GormRepository) GetRecord(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().First(out).Error
}
func (repository *GormRepository) Add(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	db1 := db.Create(out)
	return db1.Error
}

func (repository *GormRepository) Update(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	db1 := db.Update(out)
	return db1.Error
}

func (repository *GormRepository) Delete(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	db1 := db.Delete(out)
	return db1.Error
}

func (repository *GormRepository) UpdateWithMap(uow *UnitOfWork, model interface{}, value map[string]interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, value, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().Model(model).Update(value).Error
}

func DoesRecordExistForUser(db *gorm.DB, userId int, out interface{}, queryProcessors ...QueryProcessor) (bool, error) {
	if userId == 0 {
		return false, errors.New("DoesRecordExistForTenant: Invalid tenant ID")
	}
	count := 0
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return false, err
	}
	if err := db.Debug().Model(out).Where("id = ?", userId).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func DoesRecordExist(db *gorm.DB, id int, out interface{}, queryProcessors ...QueryProcessor) (bool, error) {
	if id == 0 {
		return false, errors.New("DoesRecordExistForTenant: Invalid tenant ID")
	}
	count := 0
	// fmt.Println("This....", id)
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return false, err
	}
	if err := db.Debug().Model(out).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
