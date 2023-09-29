package guru_contacts

import (
	contact_details_service "contactsoneapp/components/guru_contact_details/service"

	"github.com/google/uuid"
)

// var ContactId int = 2

type Contact struct {
	ContactId       uuid.UUID
	FirstName       string
	LastName        string
	IsActive        bool
	Contact_Details []*contact_details_service.ContactDetails
}

func NewContact(FirstName, LastName string, IsActive bool) *Contact {
	ContactDetailsTempList := make([]*contact_details_service.ContactDetails, 0)
	// contactDetailsTempItem := guru_contact_details.NewContactDetails(contactType, contactValue)
	var newObjectOfContact = &Contact{
		ContactId: uuid.New(),
		FirstName: FirstName,
		LastName:  LastName,
		IsActive:  IsActive,
		// Contact_Details: append(ContactDetailsTempList, contactDetailsTempItem),
		Contact_Details: ContactDetailsTempList,
	}

	return newObjectOfContact

}

func CreateContact(FirstName, LastName string) *Contact {
	return NewContact(FirstName, LastName, true)

}

// func (c *Contact) ReadContact() {
// 	if c.IsActive {
// 		fmt.Println("Contact Info")
// 		fmt.Printf("First Name: %s", c.FirstName)
// 		fmt.Printf("\nLast Name: %s", c.LastName)
// 		fmt.Printf("\nIsActive: %s", "True")
// 		for i := 0; i < len(c.Contact_Details); i++ {
// 			c.Contact_Details[i].ReadContactDetails()
// 		}
// 		fmt.Println()
// 	}
// }

func (c *Contact) GetContactId() uuid.UUID {
	return c.ContactId
}

// func (c *Contact) ReadContact() (readContact string) {
// 	readContact += "Contact Info" +
// 		"\n Contact Id: " + c.ContactId.String() +
// 		"\nFirst Name : " + c.FirstName +
// 		"\nLast Name : " + c.LastName +
// 		"\nIsActive : " + strconv.FormatBool(c.IsActive)
// 	for i := 0; i < len(c.Contact_Details); i++ {
// 		readContact += c.Contact_Details[i].ReadContactDetails()
// 		readContact += "\n"

// 	}

// 	return readContact

// }

func (c *Contact) ReadContact() (bool, *Contact) {
	if c.IsActive {
		return true, c
	}

	return false, c

}

func (c *Contact) DeleteContact() *Contact {
	for i := 0; i < len(c.Contact_Details); i++ {
		c.Contact_Details[i].DeleteContactDetails()
	}

	c.IsActive = false

	return c
}

// func (c *Contact) UpdateContact(updateField string, updateValue string) *Contact {
// 	switch updateField {
// 	case "FirstName":
// 		c.FirstName = updateValue
// 	case "LastName":
// 		c.LastName = updateValue
// 	case "Number", "E-Mail":
// 		c.updateContactNumberEmail(updateField, updateValue)
// 	}

// 	return c

// }

func (c *Contact) UpdateContactObject(contactTempObj *Contact) *Contact {

	if contactTempObj.FirstName != "" && contactTempObj.FirstName != c.FirstName {
		c.FirstName = contactTempObj.FirstName
	}
	if contactTempObj.LastName != "" && contactTempObj.LastName != c.LastName {
		c.LastName = contactTempObj.LastName
	}

	// switch updateField {
	// case "FirstName":
	// 	c.FirstName = updateValue
	// case "LastName":
	// 	c.LastName = updateValue
	// case "Number", "E-Mail":
	// 	c.updateContactNumberEmail(updateField, updateValue)
	// }

	return c

}

//	func (c *Contact) updateContactNumberEmail(updateField, updateValue string) {
//		for i := 0; i < len(c.Contact_Details); i++ {
//			if c.Contact_Details[i].GetType() == updateField {
//				c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
//				break
//			} else if c.Contact_Details[i].GetType() == updateField {
//				c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
//				break
//			}
//		}
//	}
func (c *Contact) GetFirstName() string {
	return c.FirstName
}
func (c *Contact) GetIsActive() bool {
	return c.IsActive
}
func (c *Contact) GetLastName() string {
	return c.LastName
}
