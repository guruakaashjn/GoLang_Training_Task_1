package guru_user

import (
	"contactsoneapp/guru_contact_details"
	"contactsoneapp/guru_contacts"
	"contactsoneapp/guru_errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

// var userId int = 1

type User struct {
	userId    uuid.UUID
	firstName string
	lastName  string
	isAdmin   bool
	isActive  bool
	Contacts  []*guru_contacts.Contact
}

var Users = make([]*User, 0)

func NewUser(firstName, lastName string, isAdmin bool) *User {

	var ContactsTempList = make([]*guru_contacts.Contact, 0)
	// var ContactsTempItem = guru_contacts.NewContact(firstName, lastName, isActive, contactType, contactValue)
	var newObjectOfUser = &User{
		userId:    uuid.New(),
		firstName: firstName,
		lastName:  lastName,
		isAdmin:   isAdmin,
		isActive:  true,
		// Contacts: append(ContactsTempList, ContactsTempItem),
		Contacts: ContactsTempList,
	}

	// fmt.Println(Users)

	return newObjectOfUser
}

// ADMIN CRUD OPERATIONS ON USERS
// ADMIN
func CreateAdmin(firstName, lastName string) *User {
	var newAdmin *User = NewUser(firstName, lastName, true)
	Users = append(Users, newAdmin)
	return newAdmin

}

func (u *User) CreateUser(firstName, lastName string) *User {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isAdmin && u.isActive {
		var newUser *User = NewUser(firstName, lastName, false)
		Users = append(Users, newUser)
		return newUser
	}
	if u.isAdmin && !u.isActive {
		panic(guru_errors.NewInvalidUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
}

func (u *User) readUser() (userInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isActive {
		userInfo +=
			"userId: " + u.userId.String() +
				"firstName: " + u.firstName +
				"lastName: " + u.lastName +
				"isAdmin: " + strconv.FormatBool(u.isAdmin) +
				"isActive: " + strconv.FormatBool(u.isActive)
		for i := 0; i < len(u.Contacts); i++ {
			userInfo += u.Contacts[i].ReadContact()
			userInfo += "\n"
		}
		return userInfo
	}

	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) ReadUserById(userIdTemp uuid.UUID) (userInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isAdmin && u.isActive {
		var requiredUser *User = GetRequiredUserObjectById(userIdTemp)

		userInfo +=
			"userId: " + requiredUser.userId.String() +
				"firstName: " + requiredUser.firstName +
				"lastName: " + requiredUser.lastName +
				"isAdmin: " + strconv.FormatBool(requiredUser.isAdmin) +
				"isActive: " + strconv.FormatBool(requiredUser.isActive)
		for i := 0; i < len(requiredUser.Contacts); i++ {
			userInfo += requiredUser.Contacts[i].ReadContact()
			userInfo += "\n"
		}
		return userInfo
	}
	if u.isAdmin && !u.isActive {
		panic(guru_errors.NewInvalidUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
}

func (u *User) ReadAllUsers() (allUserInfo string) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isAdmin && u.isActive {
		for i := 0; i < len(Users); i++ {
			allUserInfo += Users[i].readUser() + "\n"
		}
		return allUserInfo

	}
	if u.isAdmin && !u.isActive {
		panic(guru_errors.NewInvalidUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
}

func (u *User) DeleteUser(userIdTemp uuid.UUID) {
	// userIdUuid := uuid.Must(uuid.FromBytes([]byte(userId)))

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isAdmin && u.isActive {
		var requiredUser *User = GetRequiredUserObjectById(userIdTemp)

		for i := 0; i < len(requiredUser.Contacts); i++ {
			requiredUser.Contacts[i].DeleteContact()
		}

		u.isActive = false
	}
	if u.isAdmin && !u.isActive {
		panic(guru_errors.NewInvalidUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

}

func (u *User) UpdateUser(userIdTemp uuid.UUID, updateField, updateValue string) {
	// for i := 0; i < len(Users); i++ {
	// 	if Users[i].firstName == firstName {
	// 		switch updateField {
	// 		case "firstName":
	// 			u.firstName = updateValue
	// 		case "lastName":
	// 			u.lastName = updateValue
	// 		}
	// 	}
	// }

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isAdmin && u.isActive {
		var requiredUser *User = GetRequiredUserObjectById(userIdTemp)

		for i := 0; i < len(requiredUser.Contacts); i++ {
			if requiredUser.Contacts[i].GetContactId() == requiredUser.userId {
				requiredUser.Contacts[i].UpdateContact(updateField, updateValue)
				break
			}
		}

		switch updateField {
		case "firstName":
			requiredUser.firstName = updateValue
		case "lastName":
			requiredUser.lastName = updateValue
		}
	}
	if u.isAdmin && !u.isActive {
		panic(guru_errors.NewInvalidUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

}

func (u *User) CheckIsAdmin() bool {
	return u.isAdmin
}

func (u *User) GetUser() string {
	return u.firstName
}

func GetRequiredUserObjectById(userIdTemp uuid.UUID) *User {
	var requiredUser *User
	for i := 0; i < len(Users); i++ {
		if Users[i].userId == userIdTemp {
			requiredUser = Users[i]
			break
		}
	}

	return requiredUser
}

// USER CRUD OPERATIONS ON CONTACTS
// USER
func (u *User) CreateContact(firstName, lastName string, contactType, contactValue string) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		u.Contacts = append(u.Contacts, guru_contacts.CreateContact(firstName, lastName, contactType, contactValue))
		panic(guru_errors.NewContactError(guru_errors.ContactCreated).GetSpecificMessage())

	}

	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) ReadContactById(contactIdTemp uuid.UUID) (contactInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)
		contactInfo = requiredContact.ReadContact()
		panic(guru_errors.NewContactError(guru_errors.ContactRead).GetSpecificMessage())

	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())

}
func (u *User) ReadAllContact() (allContactInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if u.isActive {
		for i := 0; i < len(u.Contacts); i++ {
			allContactInfo += u.Contacts[i].ReadContact() + "\n"
		}
		panic(guru_errors.NewContactError(guru_errors.ContactReadAll).GetSpecificMessage())

	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) DeleteContact(contactIdTemp uuid.UUID) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)
		requiredContact.DeleteContact()
		panic(guru_errors.NewContactError(guru_errors.ContactDeleted).GetSpecificMessage())

	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) UpdateContact(contactIdTemp uuid.UUID, updateField, updateValue string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)

		requiredContact.UpdateContact(updateField, updateValue)

		panic(guru_errors.NewInvalidUserError(guru_errors.ContactUpdated).GetSpecificMessage())

	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) GetRequiredContactObjectById(contactIdTemp uuid.UUID) *guru_contacts.Contact {
	var requiredContact *guru_contacts.Contact
	for i := 0; i < len(u.Contacts); i++ {
		if u.Contacts[i].GetContactId() == contactIdTemp {
			requiredContact = u.Contacts[i]
		}
	}
	return requiredContact
}

// USER CRUD OPERATIONS ON CONTACT DETAILS
// USER
func (u *User) CreateContactDetails(contactIdTemp uuid.UUID, typeName, typeValue string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)
		requiredContact.Contact_Details = append(requiredContact.Contact_Details, guru_contact_details.CreateContactDetails(typeName, typeValue))
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsCreated).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())

}

func (u *User) ReadAllContactDetails(contactIdTemp uuid.UUID) (allContactDetailsInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)

		allContactDetailsInfo = requiredContact.ReadContact()

		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsReadAll).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())

}

func (u *User) ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp uuid.UUID) (contactDetailsInfo string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContactDetails *guru_contact_details.ContactDetails = u.GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp)
		contactDetailsInfo = requiredContactDetails.ReadContactDetails()
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsRead).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())

}
func (u *User) DeleteContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContactDetails *guru_contact_details.ContactDetails = u.GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp)
		requiredContactDetails.DeleteContactDetails()
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeleted).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())

}

func (u *User) UpdateContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID, keyName, keyValule string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if u.isActive {
		var requiredContactDetails *guru_contact_details.ContactDetails = u.GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp)
		requiredContactDetails.UpdateContactDetails(keyName, keyValule)
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsUpdated).GetSpecificMessage())
	}
	panic(guru_errors.NewInvalidUserError(guru_errors.UserDeleted).GetSpecificMessage())
}

func (u *User) GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp uuid.UUID) *guru_contact_details.ContactDetails {
	var requiredContactDetails *guru_contact_details.ContactDetails
	var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)
	for i := 0; i < len(requiredContact.Contact_Details); i++ {
		if requiredContact.Contact_Details[i].GetContactDetailsId() == contactDetailsIdTemp {
			requiredContactDetails = requiredContact.Contact_Details[i]
		}
	}
	return requiredContactDetails

}
