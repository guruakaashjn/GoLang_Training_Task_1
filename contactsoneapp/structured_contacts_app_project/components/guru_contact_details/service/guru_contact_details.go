package service

import (
	"contactsoneapp/models/contactinfo"
	"contactsoneapp/repository"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type ContactDetailsService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewContactInfoService(db *gorm.DB, repo repository.Repository) *ContactDetailsService {
	return &ContactDetailsService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}

func (contactDetailsService *ContactDetailsService) doesContactDetailsExist(Id uint) error {
	exists, err := repository.DoesRecordExist(contactDetailsService.db, int(Id), contactinfo.ContactInfo{}, repository.Filter("`id` = ?", Id))
	if !exists || err != nil {
		return errors.New("contact id is invalid")
	}
	return nil
}

func (contactDetailsService *ContactDetailsService) CreateContactInfo(newContactDetails *contactinfo.ContactInfo) error {
	uow := repository.NewUnitOfWork(contactDetailsService.db, false)
	defer uow.RollBack()
	err := contactDetailsService.repository.Add(uow, newContactDetails)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}
func (contactDetailsService *ContactDetailsService) GetAllContactDetails(allContactDetails *[]contactinfo.ContactInfo, contactIdTemp uint, totalCount *int, limit, offset int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(contactDetailsService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(contactDetailsService.associations, givenAssociations)

	err := contactDetailsService.repository.GetAll(uow, allContactDetails, repository.Filter("`contact_refer` =? ", contactIdTemp), repository.Paginate(limit, offset, totalCount), repository.Preload(requiredAssociations))
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (contactDetailsService *ContactDetailsService) GetContactDetailsById(requiredContactDetails *contactinfo.ContactInfo, contactIdTemp uint, idTemp int, givenAssociations []string) error {
	uow := repository.NewUnitOfWork(contactDetailsService.db, true)
	defer uow.RollBack()

	requiredAssociations := repository.FilterPreloading(contactDetailsService.associations, givenAssociations)
	err := contactDetailsService.repository.GetRecordForId(uow, uint(idTemp), requiredContactDetails, repository.Filter("`contact_refer` =?", contactIdTemp), repository.Preload(requiredAssociations))
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (contactDetailsService *ContactDetailsService) DeleteContactDetails(contactDetailsToDelete *contactinfo.ContactInfo) error {
	err := contactDetailsService.doesContactDetailsExist(contactDetailsToDelete.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(contactDetailsService.db, false)
	defer uow.RollBack()
	if err := contactDetailsService.repository.UpdateWithMap(uow, contactDetailsToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
	}, repository.Filter("`id`=?", contactDetailsToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (contactDetailsService *ContactDetailsService) UpdateContactDetails(contactDetailsToUpdate *contactinfo.ContactInfo) error {
	err := contactDetailsService.doesContactDetailsExist(contactDetailsToUpdate.ID)
	if err != nil {
		return nil
	}
	uow := repository.NewUnitOfWork(contactDetailsService.db, false)
	defer uow.RollBack()
	tempContactDetails := contactinfo.ContactInfo{}
	err = contactDetailsService.repository.GetRecordForId(uow, contactDetailsToUpdate.ID, &tempContactDetails, repository.Select("`created_at`"), repository.Filter("`id` = ?", contactDetailsToUpdate.ID))
	if err != nil {
		return nil
	}
	contactDetailsToUpdate.CreatedAt = tempContactDetails.CreatedAt
	err = contactDetailsService.repository.Save(uow, contactDetailsToUpdate)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

// package service

// import (
// 	"github.com/google/uuid"
// )

// // var ContactDetailsId = 3

// type ContactDetails struct {
// 	ContactDetailsId uuid.UUID
// 	TypeName         string
// 	TypeValue        string
// 	IsActive         bool
// }

// func NewContactDetails(TypeName string, TypeValue string) *ContactDetails {

// 	var newObjectOfContactDetails = &ContactDetails{
// 		ContactDetailsId: uuid.New(),
// 		TypeName:         TypeName,
// 		TypeValue:        TypeValue,
// 		IsActive:         true,
// 	}
// 	return newObjectOfContactDetails
// }

// func CreateContactDetails(TypeName, TypeValue string) *ContactDetails {
// 	return NewContactDetails(TypeName, TypeValue)
// }

// func (cd *ContactDetails) UpdateContactDetails(keyName string, keyValue string) *ContactDetails {
// 	cd.TypeName = keyName
// 	cd.TypeValue = keyValue
// 	return cd
// }

// func (cd *ContactDetails) UpdateContactDetailsObject(contactDetailsTempObj *ContactDetails) *ContactDetails {
// 	if contactDetailsTempObj.TypeName != "" && contactDetailsTempObj.TypeName != cd.TypeName {
// 		cd.TypeName = contactDetailsTempObj.TypeName
// 	}
// 	if contactDetailsTempObj.TypeValue != "" && contactDetailsTempObj.TypeValue != cd.TypeValue {
// 		cd.TypeValue = contactDetailsTempObj.TypeValue
// 	}
// 	return cd
// }

// // func (cd *ContactDetails) ReadContactDetails() {
// // 	fmt.Printf("\nContact Details")
// // 	fmt.Printf("\nType : %s and Type Value : %s", cd.TypeName, cd.TypeValue)
// // }

// // func (cd *ContactDetails) ReadContactDetails() (readContactDetails string) {
// // 	readContactDetails += "Contact Details" +
// // 		"\nContact Details Id: " + cd.ContactDetailsId.String() +
// // 		"\nType : " + cd.TypeName +
// // 		" and Type Value : " + cd.TypeValue +
// // 		"\nIsActive : " + strconv.FormatBool(cd.IsActive)
// // 	return readContactDetails
// // }

// func (cd *ContactDetails) ReadContactDetails() (bool, *ContactDetails) {
// 	if cd.IsActive {
// 		return true, cd
// 	}
// 	return false, cd
// }

// func (cd *ContactDetails) DeleteContactDetails() *ContactDetails {
// 	// cd.TypeName = ""
// 	// cd.TypeValue = ""
// 	cd.IsActive = false
// 	return cd
// }

// func (cd *ContactDetails) GetType() string {
// 	return cd.TypeName
// }
// func (cd *ContactDetails) GetTypeValue() string {
// 	return cd.TypeValue
// }

// func (cd *ContactDetails) GetContactDetailsId() uuid.UUID {
// 	return cd.ContactDetailsId
// }
