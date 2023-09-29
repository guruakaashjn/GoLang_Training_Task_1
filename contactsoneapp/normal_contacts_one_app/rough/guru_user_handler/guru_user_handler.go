package guru_user_handler

import (
	"contactsoneapp/guru_contact_details"
	"contactsoneapp/guru_contacts"
	"contactsoneapp/guru_errors"
	"contactsoneapp/guru_user"
	"fmt"
)

func UserHandler() (check bool, firstName string) {
	fmt.Println("Enter F_Name : ")
	// var firstName string
	fmt.Scan(&firstName)
	defer func() {
		if a := recover(); a != nil {
			fmt.Println("Recovered", a)
			check = false
		}

	}()
	check = userCheck(firstName)
	return check, firstName
}

func userCheck(firstName string) bool {
	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName && !guru_user.Users[i].CheckIsAdmin() {
			return true
		}
		if guru_user.Users[i].GetUser() == firstName {
			msg := guru_user.Users[i].GetUser() + " is an Admin!"
			panic(guru_errors.NewInvalidUserError(msg))
		}
	}
	msg := firstName + " does not exist!"
	panic(guru_errors.NewInvalidUserError(msg))

}

func UserPrivilages(firstName string) {
	for i := 0; i < 1; {
		// var userObj1 = guru_user.GetUserObject(firstName)
		fmt.Println("Menu:")
		fmt.Println("1. Create a Contact")
		fmt.Println("2. Read all Contacts")
		fmt.Println("3. Update a Contact")
		fmt.Println("4. Add Contact Details to existing Contact")
		fmt.Println("5. Delete a Contact")
		fmt.Println("6. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var contactObj1 *guru_contacts.Contact
			contactObj1 = getInputFromUser()
			// newContacts := append(userObj1.Contacts, contactObj1)
			for i := 0; i < len(guru_user.Users); i++ {
				if guru_user.Users[i].GetUser() == firstName {
					guru_user.Users[i].Contacts = append(guru_user.Users[i].Contacts, contactObj1)
				}
			}
		case 2:
			readAllContacts(firstName)
		case 3:
			updateContact(firstName)
		case 4:
			addContactDetailsToExistingContact(firstName)
		case 5:
			deleteContact(firstName)
		case 6:
			i++

		}
	}
}

func addContactDetailsToExistingContact(firstName string) {
	fmt.Println("Enter contact first name: ")
	var fName string
	fmt.Scan(&fName)

	fmt.Println("Enter contact details type name: ")
	var typeName string
	fmt.Scan(&typeName)
	fmt.Println("Enter contact details type value: ")
	var typeValue string
	fmt.Scan(&typeValue)

	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			for j := 0; j < len(guru_user.Users[i].Contacts); j++ {
				if guru_user.Users[i].Contacts[j].GetFirstName() == fName {
					guru_user.Users[i].Contacts[j].Contact_Details = append(guru_user.Users[i].Contacts[j].Contact_Details, guru_contact_details.NewContactDetails(typeName, typeValue))
					break
				}
			}
		}
	}
}

func deleteContact(firstName string) {

	fmt.Println("Enter Contact First Name to be deleted: ")
	var fName string
	fmt.Scan(&fName)
	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			for j := 0; j < len(guru_user.Users[i].Contacts); j++ {
				if guru_user.Users[i].Contacts[j].GetFirstName() == fName {
					guru_user.Users[i].Contacts[j].DeleteContact()
				}
			}
		}
	}

}

func updateContact(firstName string) {
	fmt.Println("Enter Contact FirstName to update: ")
	var fName string
	fmt.Scan(&fName)
	fmt.Println("Enter Contact update field name: ")
	var updateField string
	fmt.Scan(&updateField)
	fmt.Println("Enter Contact update field value: ")
	var updateValue string
	fmt.Scan(&updateValue)

	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			for j := 0; j < len(guru_user.Users[i].Contacts); j++ {
				if guru_user.Users[i].Contacts[j].GetFirstName() == fName {
					guru_user.Users[i].Contacts[j].UpdateContact(updateField, updateValue)
					break
				}
			}
			break

		}
	}
}

func readAllContacts(firstName string) {
	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			for j := 0; j < len(guru_user.Users[i].Contacts); j++ {
				guru_user.Users[i].Contacts[j].ReadContact()
			}
			break
		}
	}
}

func getInputFromUser() *guru_contacts.Contact {
	var firstName string
	var lastName string
	var isActive bool

	var contactChoice int
	var contactType string
	var contactValue string

	fmt.Println("Enter firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter contact type: ")
	fmt.Println("1. Number")
	fmt.Println("2. E-Mail")
	fmt.Scan(&contactChoice)
	switch contactChoice {
	case 1:
		contactType = "Number"
	case 2:
		contactType = "E-Mail"
	}

	fmt.Println("Enter contact value: ")
	fmt.Scan(&contactValue)

	isActive = true

	contactObj1 := guru_contacts.NewContact(firstName, lastName, isActive, contactType, contactValue)
	return contactObj1
}
