package service

import (
	"contactsoneapp/models/contact"
	"contactsoneapp/repository"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type ContactService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewContactService(db *gorm.DB, repo repository.Repository) *ContactService {
	return &ContactService{
		db:           db,
		repository:   repo,
		associations: []string{"ContactInfos"},
	}
}

func (contactService *ContactService) doesContactExist(Id uint) error {
	exists, err := repository.DoesRecordExist(contactService.db, int(Id), contact.Contact{}, repository.Filter("`id` = ?", Id))

	if !exists || err != nil {
		return errors.New("contact id is invalid")
	}
	return nil
}
func (contactService *ContactService) CreateContact(newContact *contact.Contact) error {
	uow := repository.NewUnitOfWork(contactService.db, false)
	defer uow.RollBack()
	err := contactService.repository.Add(uow, newContact)
	if err != nil {
		uow.RollBack()
		return err

	}
	uow.Commit()
	return nil
}
func (contactService *ContactService) GetAllContacts(allContacts *[]contact.Contact, totalCount *int) error {
	uow := repository.NewUnitOfWork(contactService.db, true)
	defer uow.RollBack()
	err := contactService.repository.GetAll(uow, allContacts, contactService.associations)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (contactService *ContactService) GetContactById(requiredContact *contact.Contact, idTemp int) error {
	uow := repository.NewUnitOfWork(contactService.db, true)
	defer uow.RollBack()

	err := contactService.repository.GetRecordForId(uow, uint(idTemp), requiredContact)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (contactService *ContactService) DeleteContact(contactToDelete *contact.Contact) error {
	err := contactService.doesContactExist(contactToDelete.ID)
	// fmt.Println(contactToDelete.ID)
	// fmt.Println("A: ", contactToDelete.ID)
	if err != nil {
		// fmt.Println("This ....", err)
		return err
	}
	uow := repository.NewUnitOfWork(contactService.db, false)

	defer uow.RollBack()
	if err := contactService.repository.UpdateWithMap(uow, contactToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
	},
		repository.Filter("`id`=?", contactToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (contactService *ContactService) UpdateContact(contactToUpdate *contact.Contact) error {
	err := contactService.doesContactExist(contactToUpdate.ID)
	if err != nil {
		return nil
	}
	// fmt.Println("A: ", contactToUpdate.ID)
	uow := repository.NewUnitOfWork(contactService.db, false)
	defer uow.RollBack()
	tempContact := contact.Contact{}

	err = contactService.repository.GetRecordForId(uow, contactToUpdate.ID, &tempContact, repository.Select("`created_at`"), repository.Filter("`id` = ?", contactToUpdate.ID))
	if err != nil {
		return nil
	}
	contactToUpdate.CreatedAt = tempContact.CreatedAt
	err = contactService.repository.Save(uow, contactToUpdate)
	if err != nil {
		return err

	}
	uow.Commit()
	return nil
}

// package service

// import (
// 	contact_details_service "contactsoneapp/components/guru_contact_details/service"

// 	"github.com/google/uuid"
// )

// // var ContactId int = 2

// type Contact struct {
// 	ContactId       uuid.UUID
// 	FirstName       string
// 	LastName        string
// 	IsActive        bool
// 	Contact_Details []*contact_details_service.ContactDetails
// }

// func NewContact(FirstName, LastName string, IsActive bool) *Contact {
// 	ContactDetailsTempList := make([]*contact_details_service.ContactDetails, 0)
// 	// contactDetailsTempItem := guru_contact_details.NewContactDetails(contactType, contactValue)
// 	var newObjectOfContact = &Contact{
// 		ContactId: uuid.New(),
// 		FirstName: FirstName,
// 		LastName:  LastName,
// 		IsActive:  IsActive,
// 		// Contact_Details: append(ContactDetailsTempList, contactDetailsTempItem),
// 		Contact_Details: ContactDetailsTempList,
// 	}

// 	return newObjectOfContact

// }

// func CreateContact(FirstName, LastName string) *Contact {
// 	return NewContact(FirstName, LastName, true)

// }

// // func (c *Contact) ReadContact() {
// // 	if c.IsActive {
// // 		fmt.Println("Contact Info")
// // 		fmt.Printf("First Name: %s", c.FirstName)
// // 		fmt.Printf("\nLast Name: %s", c.LastName)
// // 		fmt.Printf("\nIsActive: %s", "True")
// // 		for i := 0; i < len(c.Contact_Details); i++ {
// // 			c.Contact_Details[i].ReadContactDetails()
// // 		}
// // 		fmt.Println()
// // 	}
// // }

// func (c *Contact) GetContactId() uuid.UUID {
// 	return c.ContactId
// }

// // func (c *Contact) ReadContact() (readContact string) {
// // 	readContact += "Contact Info" +
// // 		"\n Contact Id: " + c.ContactId.String() +
// // 		"\nFirst Name : " + c.FirstName +
// // 		"\nLast Name : " + c.LastName +
// // 		"\nIsActive : " + strconv.FormatBool(c.IsActive)
// // 	for i := 0; i < len(c.Contact_Details); i++ {
// // 		readContact += c.Contact_Details[i].ReadContactDetails()
// // 		readContact += "\n"

// // 	}

// // 	return readContact

// // }

// func (c *Contact) ReadContact() (bool, *Contact) {
// 	if c.IsActive {
// 		return true, c
// 	}

// 	return false, c

// }

// func (c *Contact) DeleteContact() *Contact {
// 	for i := 0; i < len(c.Contact_Details); i++ {
// 		c.Contact_Details[i].DeleteContactDetails()
// 	}

// 	c.IsActive = false

// 	return c
// }

// // func (c *Contact) UpdateContact(updateField string, updateValue string) *Contact {
// // 	switch updateField {
// // 	case "FirstName":
// // 		c.FirstName = updateValue
// // 	case "LastName":
// // 		c.LastName = updateValue
// // 	case "Number", "E-Mail":
// // 		c.updateContactNumberEmail(updateField, updateValue)
// // 	}

// // 	return c

// // }

// func (c *Contact) UpdateContactObject(contactTempObj *Contact) *Contact {

// 	if contactTempObj.FirstName != "" && contactTempObj.FirstName != c.FirstName {
// 		c.FirstName = contactTempObj.FirstName
// 	}
// 	if contactTempObj.LastName != "" && contactTempObj.LastName != c.LastName {
// 		c.LastName = contactTempObj.LastName
// 	}

// 	// switch updateField {
// 	// case "FirstName":
// 	// 	c.FirstName = updateValue
// 	// case "LastName":
// 	// 	c.LastName = updateValue
// 	// case "Number", "E-Mail":
// 	// 	c.updateContactNumberEmail(updateField, updateValue)
// 	// }

// 	return c

// }

// //	func (c *Contact) updateContactNumberEmail(updateField, updateValue string) {
// //		for i := 0; i < len(c.Contact_Details); i++ {
// //			if c.Contact_Details[i].GetType() == updateField {
// //				c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
// //				break
// //			} else if c.Contact_Details[i].GetType() == updateField {
// //				c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
// //				break
// //			}
// //		}
// //	}
// func (c *Contact) GetFirstName() string {
// 	return c.FirstName
// }
// func (c *Contact) GetIsActive() bool {
// 	return c.IsActive
// }
// func (c *Contact) GetLastName() string {
// 	return c.LastName
// }
