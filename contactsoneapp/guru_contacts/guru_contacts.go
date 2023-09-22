package guru_contacts

import (
	"contactsoneapp/guru_contact_details"

	"github.com/google/uuid"
)

// var ContactId int = 2

type Contact struct {
	contactId       uuid.UUID
	firstName       string
	lastName        string
	isActive        bool
	Contact_Details []*guru_contact_details.ContactDetails
}

func NewContact(firstName, lastName string, isActive bool, contactType, contactValue string) *Contact {
	ContactDetailsTempList := make([]*guru_contact_details.ContactDetails, 0)
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

func CreateContact(firstName, lastName string, contactType, contactValue string) *Contact {
	return NewContact(firstName, lastName, true, contactType, contactValue)

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

func (c *Contact) ReadContact() *Contact {

	return c

}

func (c *Contact) DeleteContact() {
	for i := 0; i < len(c.Contact_Details); i++ {
		c.Contact_Details[i].DeleteContactDetails()
	}

	c.isActive = false
}

func (c *Contact) UpdateContact(updateField string, updateValue string) {
	switch updateField {
	case "firstName":
		c.firstName = updateValue
	case "lastName":
		c.lastName = updateValue
	case "Number", "E-Mail":
		c.updateContactNumberEmail(updateField, updateValue)
	}

}

func (c *Contact) updateContactNumberEmail(updateField, updateValue string) {
	for i := 0; i < len(c.Contact_Details); i++ {
		if c.Contact_Details[i].GetType() == updateField {
			c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
			break
		} else if c.Contact_Details[i].GetType() == updateField {
			c.Contact_Details[i].UpdateContactDetails(updateField, updateValue)
			break
		}
	}
}
func (c *Contact) GetFirstName() string {
	return c.firstName
}

func (c *Contact) AddContactDetailsToExistingContact() {

}

// func createAndAddRecord
