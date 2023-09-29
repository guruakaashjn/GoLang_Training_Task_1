package guru_contacts

import (
	contact_details_service "contactsoneapp/components/guru_contact_details/service"

	"github.com/google/uuid"
)

// var ContactId int = 2

type Contact struct {
	contactId       uuid.UUID
	firstName       string
	lastName        string
	isActive        bool
	Contact_Details []*contact_details_service.ContactDetails
}

func NewContact(firstName, lastName string, isActive bool) *Contact {
	ContactDetailsTempList := make([]*contact_details_service.ContactDetails, 0)
	// contactDetailsTempItem := guru_contact_details.NewContactDetails(contactType, contactValue)
	var newObjectOfContact = &Contact{
		contactId: uuid.New(),
		firstName: firstName,
		lastName:  lastName,
		isActive:  isActive,
		// Contact_Details: append(ContactDetailsTempList, contactDetailsTempItem),
		Contact_Details: ContactDetailsTempList,
	}

	return newObjectOfContact

}

func CreateContact(firstName, lastName string) *Contact {
	return NewContact(firstName, lastName, true)

}

// func (c *Contact) ReadContact() {
// 	if c.isActive {
// 		fmt.Println("Contact Info")
// 		fmt.Printf("First Name: %s", c.firstName)
// 		fmt.Printf("\nLast Name: %s", c.lastName)
// 		fmt.Printf("\nisActive: %s", "True")
// 		for i := 0; i < len(c.Contact_Details); i++ {
// 			c.Contact_Details[i].ReadContactDetails()
// 		}
// 		fmt.Println()
// 	}
// }

func (c *Contact) GetContactId() uuid.UUID {
	return c.contactId
}

// func (c *Contact) ReadContact() (readContact string) {
// 	readContact += "Contact Info" +
// 		"\n Contact Id: " + c.contactId.String() +
// 		"\nFirst Name : " + c.firstName +
// 		"\nLast Name : " + c.lastName +
// 		"\nisActive : " + strconv.FormatBool(c.isActive)
// 	for i := 0; i < len(c.Contact_Details); i++ {
// 		readContact += c.Contact_Details[i].ReadContactDetails()
// 		readContact += "\n"

// 	}

// 	return readContact

// }

func (c *Contact) ReadContact() (bool, *Contact) {
	if c.isActive {
		return true, c
	}

	return false, c

}

func (c *Contact) DeleteContact() *Contact {
	for i := 0; i < len(c.Contact_Details); i++ {
		c.Contact_Details[i].DeleteContactDetails()
	}

	c.isActive = false

	return c
}

// func (c *Contact) UpdateContact(updateField string, updateValue string) *Contact {
// 	switch updateField {
// 	case "firstName":
// 		c.firstName = updateValue
// 	case "lastName":
// 		c.lastName = updateValue
// 	case "Number", "E-Mail":
// 		c.updateContactNumberEmail(updateField, updateValue)
// 	}

// 	return c

// }

func (c *Contact) UpdateContactObject(contactTempObj *Contact) *Contact {

	if contactTempObj.firstName != "" && contactTempObj.firstName != c.firstName {
		c.firstName = contactTempObj.firstName
	}
	if contactTempObj.lastName != "" && contactTempObj.lastName != c.lastName {
		c.lastName = contactTempObj.lastName
	}

	// switch updateField {
	// case "firstName":
	// 	c.firstName = updateValue
	// case "lastName":
	// 	c.lastName = updateValue
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
	return c.firstName
}
func (c *Contact) GetIsActive() bool {
	return c.isActive
}
func (c *Contact) GetLastName() string {
	return c.lastName
}
