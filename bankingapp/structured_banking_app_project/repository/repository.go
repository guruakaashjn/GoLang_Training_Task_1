package repository

import (
	"bankingapp/errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Add(uow *UnitOfWork, out interface{}) error
	GetAll(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	GetRecordForId(uow *UnitOfWork, id uint, out interface{}, queryProcessors ...QueryProcessor) error
	GetRecord(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	Update(uow *UnitOfWork, out interface{}) error
	UpdateWithMap(uow *UnitOfWork, model interface{}, value map[string]interface{}, queryProcessors ...QueryProcessor) error
	Delete(uow *UnitOfWork, out interface{}) error
	Save(uow *UnitOfWork, value interface{}) error
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

// GormRepository implements Repository interface if it implements all of the functions described in Repository interface.

func (repository *GormRepository) Save(uow *UnitOfWork, value interface{}) error {
	return uow.DB.Debug().Save(value).Error
}

func (repository *GormRepository) Add(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	return db.Create(out).Error
}

func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return err
	}
	db = db.Debug()
	// for _, association := range associations {
	// 	db = db.Preload(association)
	// }
	return db.Find(out).Error
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
	db = db.Debug()
	// for _, association := range associations {
	// 	db = db.Preload(association)
	// }
	return db.First(out).Error
}

func (repository *GormRepository) Update(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	return db.Model(out).Update(out).Error
}
func (repository *GormRepository) UpdateWithMap(uow *UnitOfWork, model interface{}, value map[string]interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, value, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().Model(model).Update(value).Error
}

func (repository *GormRepository) Delete(uow *UnitOfWork, out interface{}) error {
	db := uow.DB
	return db.Delete(out).Error
}

// Helper Functions

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
func Join(query string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Joins(query, args...)
		return db, nil
	}
}
func Scan(out interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Scan(out)
		return db, nil
	}
}
func GroupBy(groupstr ...string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		for _, entity := range groupstr {
			db = db.Group(entity)
		}
		return db, nil
	}
}
func OrderBy(entity string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Order(entity)
		return db, nil
	}
}
func Distinct(columns ...string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		length := len(columns)
		if length == 0 {
			return db, nil
		}
		str := "DISTINCT "
		for index, column := range columns {
			if length-1 != index {
				str += column + ","
			} else {
				str += column
			}
		}
		db = db.Select(str)
		return db, nil
	}
}
func RawQuery(sql string, values ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Raw(sql, values...)
		return db, db.Error
	}
}
func Table(tableName string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Table(tableName)
		return db, nil
	}
}

func Offset(offset interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Offset(offset)
		return db, nil
	}
}

func Limit(limit interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Limit(limit)
		return db, nil
	}
}

func SearchFilter(searchParams map[string]interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		for field, value := range searchParams {
			db = db.Debug().Where("`"+field+"`"+"LIKE ?", "%"+value.(string)+"%")
		}
		return db, nil
	}
}
func FilterWithOperator(columnNames []string, conditions []string, operators []string, values []interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {

		if len(columnNames) != len(conditions) && len(conditions) != len(values) {
			return db, nil
		}

		if len(conditions) == 1 {
			if values[0] == nil {
				db = db.Where(fmt.Sprintf("%v %v", columnNames[0], conditions[0]))
				return db, nil
			}
			db = db.Where(fmt.Sprintf("%v %v", columnNames[0], conditions[0]), values[0])
			return db, nil
		}
		if len(columnNames)-1 != len(operators) {
			return db, nil
		}

		str := ""
		nums := []int{}
		for index := 0; index < len(columnNames); index++ {
			if values[index] == nil {
				nums = append(nums, index)
			}
			if index == len(columnNames)-1 {
				str = fmt.Sprintf("%v%v %v", str, columnNames[index], conditions[index])
			} else {
				str = fmt.Sprintf("%v%v %v %v ", str, columnNames[index], conditions[index], operators[index])
			}
		}
		for ind, num := range nums {
			values = append(values[:num], values[num+1:]...)
			for i := ind; i < len(nums); i++ {
				// This is done to adjust indexes because we sliced.
				nums[i] = nums[i] - 1
			}
		}
		db = db.Where(str, values...)
		return db, nil
	}
}

func FilterPreloading(availableAssociations, givenAssociations []string) (requiredAssociations []string) {
	for _, association := range givenAssociations {
		for _, availableAssociation := range availableAssociations {
			if association == availableAssociation {
				requiredAssociations = append(requiredAssociations, association)
			}
		}
	}

	return requiredAssociations

}

func Preload(preloads interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {

		for _, preload := range preloads.([]string) {
			// fmt.Println(preload)

			db = db.Preload(preload)
		}

		return db, nil
	}
}

func Paginate(limit, offset int, totalCount *int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if out != nil {
			if totalCount != nil {
				if err := db.Model(out).Count(totalCount).Error; err != nil {
					return db, err
				}
			}
		}
		if limit != -1 {
			db = db.Limit(limit)
		}
		if offset > 0 {
			db = db.Offset(limit * offset)
		}
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

func DoesRecordExist(db *gorm.DB, id int, out interface{}, queryProcessors ...QueryProcessor) (bool, error) {
	if id == 0 {
		return false, errors.NewValidationError("DoesRecordExist: Invalid id")
	}
	count := 0
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
