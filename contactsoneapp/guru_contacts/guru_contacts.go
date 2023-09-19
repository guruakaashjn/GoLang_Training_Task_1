package guru_contacts

import (
	"contactsoneapp/guru_contact_details"
	"fmt"

	"github.com/google/uuid"
)

var ContactId int = 2

type Contact struct {
	contact_ID      uuid.UUID
	f_name          string
	l_name          string
	isActive        bool
	Contact_Details []*guru_contact_details.ContactDetails
}

func NewContact(firstName, lastName string, isActive bool, contactType, contactValue string) *Contact {
	ContactDetailsTempList := make([]*guru_contact_details.ContactDetails, 0)
	contactDetailsTempItem := guru_contact_details.NewContactDetails(contactType, contactValue)
	var newObjectOfContact = &Contact{
		contact_ID:      uuid.New(),
		f_name:          firstName,
		l_name:          lastName,
		isActive:        isActive,
		Contact_Details: append(ContactDetailsTempList, contactDetailsTempItem),
	}

	return newObjectOfContact

}

func (c *Contact) ReadContact() {
	if c.isActive {
		fmt.Println("Contact Info")
		fmt.Printf("First Name: %s", c.f_name)
		fmt.Printf("\nLast Name: %s", c.l_name)
		fmt.Printf("\nisActive: %s", "True")
		for i := 0; i < len(c.Contact_Details); i++ {
			c.Contact_Details[i].ReadContactDetails()
		}
	}
}

func (c *Contact) DeleteContact() {
	c.isActive = false
}

func (c *Contact) UpdateContact(updateField string, updateValue string) {
	switch updateField {
	case "F_Name":
		c.f_name = updateValue
	case "L_Name":
		c.l_name = updateValue
	case "Number", "E-Mail":
		c.UpdateContactNumberEmail(updateField, updateValue)
	}

}

func (c *Contact) UpdateContactNumberEmail(updateField, updateValue string) {
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
	return c.f_name
}

// func createAndAddRecord
