package guru_user

import (
	"contactsoneapp/guru_contact_details"
	"contactsoneapp/guru_contacts"
	"contactsoneapp/guru_errors"
	"fmt"

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

	var ContactsTempList []*guru_contacts.Contact = make([]*guru_contacts.Contact, 0)
	// var ContactsTempItem = guru_contacts.NewContact(firstName, lastName, true, contactType, contactValue)

	var newObjectOfUser = &User{
		userId:    uuid.New(),
		firstName: firstName,
		lastName:  lastName,
		isAdmin:   isAdmin,
		isActive:  true,
		// Contacts: append(ContactsTempList, ContactsTempItem),
		Contacts: ContactsTempList,
	}
	Users = append(Users, newObjectOfUser)

	// fmt.Println(Users)

	return newObjectOfUser
}

// ADMIN CRUD OPERATIONS ON USERS
// ADMIN
func CreateAdmin(firstName, lastName string) (adminUser *User) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	adminUser = NewUser(firstName, lastName, true)
	panic(guru_errors.NewUserError(guru_errors.AdminCreated).GetSpecificMessage())

}

func (u *User) CreateUser(firstName, lastName string) (newUser *User) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isAdmin {
		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
	}
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}

	newUser = NewUser(firstName, lastName, false)
	panic(guru_errors.NewUserError(guru_errors.UserCreated).GetSpecificMessage())

}

// func (u *User) readUser() (userInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if u.isActive {
// 		userInfo +=
// 			"userId: " + u.userId.String() +
// 				"firstName: " + u.firstName +
// 				"lastName: " + u.lastName +
// 				"isAdmin: " + strconv.FormatBool(u.isAdmin) +
// 				"isActive: " + strconv.FormatBool(u.isActive)
// 		for i := 0; i < len(u.Contacts); i++ {
// 			userInfo += u.Contacts[i].ReadContact()
// 			userInfo += "\n"
// 		}
// 		return userInfo
// 	}

// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())
// }

// func (u *User) ReadUserById(userIdTemp uuid.UUID) (userInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if u.isAdmin && u.isActive {
// 		var requiredUser *User = GetRequiredUserObjectById(userIdTemp)

// 		userInfo +=
// 			"userId: " + requiredUser.userId.String() +
// 				"firstName: " + requiredUser.firstName +
// 				"lastName: " + requiredUser.lastName +
// 				"isAdmin: " + strconv.FormatBool(requiredUser.isAdmin) +
// 				"isActive: " + strconv.FormatBool(requiredUser.isActive)
// 		for i := 0; i < len(requiredUser.Contacts); i++ {
// 			userInfo += requiredUser.Contacts[i].ReadContact()
// 			userInfo += "\n"
// 		}
// 		return userInfo
// 	}
// 	if u.isAdmin && !u.isActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// }

// func (u *User) ReadAllUsers() (allUserInfo string) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if u.isAdmin && u.isActive {
// 		for i := 0; i < len(Users); i++ {
// 			allUserInfo += Users[i].readUser() + "\n"
// 		}
// 		return allUserInfo

// 	}
// 	if u.isAdmin && !u.isActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// }

func (u *User) ReadUserById(userIdTemp uuid.UUID) (userInfo *User) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !u.isAdmin {
		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

	}
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}

	var requiredUser *User
	for i := 0; i < len(Users); i++ {
		if Users[i].userId == userIdTemp {
			requiredUser = Users[i]
			break
		}
	}
	if !requiredUser.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

	}
	userInfo = requiredUser
	panic(guru_errors.NewUserError(guru_errors.UserRead).GetSpecificMessage())

}

func (u *User) ReadAllUsers() (allUserInfo []*User) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isAdmin {
		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

	}

	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}

	for i := 0; i < len(Users); i++ {
		if Users[i].isActive {
			allUserInfo = append(allUserInfo, Users[i])
		}

	}
	panic(guru_errors.NewUserError(guru_errors.UserReadAll).GetSpecificMessage())

}

func (u *User) DeleteUser(userIdTemp uuid.UUID) (deletedUser *User) {
	// userIdUuid := uuid.Must(uuid.FromBytes([]byte(userId)))

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isAdmin {
		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
	}
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}

	requiredUser := u.ReadUserById(userIdTemp)
	if requiredUser == nil {
		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

	}

	for i := 0; i < len(requiredUser.Contacts); i++ {
		requiredUser.Contacts[i].DeleteContact()
	}

	requiredUser.isActive = false
	deletedUser = requiredUser
	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

}

func (u *User) UpdateUser(userIdTemp uuid.UUID, updateField, updateValue string) (updatedUser *User) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !u.isAdmin {
		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
	}
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
	}

	requiredUser := u.ReadUserById(userIdTemp)
	if requiredUser == nil {
		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())
	}

	// for i := 0; i < len(requiredUser.Contacts); i++ {
	// 	if requiredUser.Contacts[i].GetContactId() == requiredUser.userId {
	// 		requiredUser.Contacts[i].UpdateContact(updateField, updateValue)
	// 		break
	// 	}
	// }

	switch updateField {
	case "firstName":
		requiredUser.firstName = updateValue
	case "lastName":
		requiredUser.lastName = updateValue
	}
	updatedUser = requiredUser
	panic(guru_errors.NewUserError(guru_errors.UserUpdated).GetSpecificMessage())

}

func (u *User) CheckIsAdmin() bool {
	return u.isAdmin
}

func (u *User) GetUser() string {
	return u.firstName
}

// USER CRUD OPERATIONS ON CONTACTS
// USER
func (u *User) CreateContact(firstName, lastName string) (newContact *guru_contacts.Contact) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	newContact = guru_contacts.CreateContact(firstName, lastName)
	u.Contacts = append(u.Contacts, newContact)
	panic(guru_errors.NewContactError(guru_errors.ContactCreated).GetSpecificMessage())

}

// func (u *User) ReadContactById(contactIdTemp uuid.UUID) (contactInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if u.isActive {
// 		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)
// 		contactInfo = requiredContact.ReadContact()
// 		panic(guru_errors.NewContactError(guru_errors.ContactRead).GetSpecificMessage())

// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// }
// func (u *User) ReadAllContact() (allContactInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if u.isActive {
// 		for i := 0; i < len(u.Contacts); i++ {
// 			allContactInfo += u.Contacts[i].ReadContact() + "\n"
// 		}
// 		panic(guru_errors.NewContactError(guru_errors.ContactReadAll).GetSpecificMessage())

// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())
// }

func (u *User) ReadContactById(contactIdTemp uuid.UUID) (contactInfo *guru_contacts.Contact) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	var requiredContact *guru_contacts.Contact

	for i := 0; i < len(u.Contacts); i++ {
		if u.Contacts[i].GetContactId() == contactIdTemp {
			requiredContact = u.Contacts[i]
			break
		}
	}

	if !requiredContact.GetIsActive() {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())

	}
	contactInfo = requiredContact
	panic(guru_errors.NewContactError(guru_errors.ContactRead).GetSpecificMessage())

}
func (u *User) ReadAllContact() (allContactInfo []*guru_contacts.Contact) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if !u.isActive {

		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	for i := 0; i < len(u.Contacts); i++ {
		flag, contactTemp := u.Contacts[i].ReadContact()
		if flag {
			allContactInfo = append(allContactInfo, contactTemp)
		}

	}
	panic(guru_errors.NewContactError(guru_errors.ContactReadAll).GetSpecificMessage())

}

func (u *User) DeleteContact(contactIdTemp uuid.UUID) (deletedContact *guru_contacts.Contact) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	var requiredContact *guru_contacts.Contact = u.ReadContactById(contactIdTemp)
	if requiredContact == nil {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
	}

	deletedContact = requiredContact.DeleteContact()
	panic(guru_errors.NewContactError(guru_errors.ContactDeleted).GetSpecificMessage())

}

func (u *User) UpdateContact(contactIdTemp uuid.UUID, updateField, updateValue string) (updatedContact *guru_contacts.Contact) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}

	var requiredContact *guru_contacts.Contact = u.ReadContactById(contactIdTemp)
	if requiredContact == nil {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
	}

	updatedContact = requiredContact.UpdateContact(updateField, updateValue)
	panic(guru_errors.NewUserError(guru_errors.ContactUpdated).GetSpecificMessage())

}

// USER CRUD OPERATIONS ON CONTACT DETAILS
// USER
func (u *User) CreateContactDetails(contactIdTemp uuid.UUID, typeName, typeValue string) (newContactDetails *guru_contact_details.ContactDetails) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	var requiredContact *guru_contacts.Contact = u.ReadContactById(contactIdTemp)
	if requiredContact == nil {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
	}

	newContactDetails = guru_contact_details.CreateContactDetails(typeName, typeValue)
	requiredContact.Contact_Details = append(requiredContact.Contact_Details, newContactDetails)
	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsCreated).GetSpecificMessage())

}

// func (u *User) ReadAllContactDetails(contactIdTemp uuid.UUID) (allContactDetailsInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if u.isActive {
// 		var requiredContact *guru_contacts.Contact = u.GetRequiredContactObjectById(contactIdTemp)

// 		allContactDetailsInfo = requiredContact.ReadContact()

// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsReadAll).GetSpecificMessage())
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// }

// func (u *User) ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp uuid.UUID) (contactDetailsInfo string) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if u.isActive {
// 		var requiredContactDetails *guru_contact_details.ContactDetails = u.GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp)
// 		contactDetailsInfo = requiredContactDetails.ReadContactDetails()
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsRead).GetSpecificMessage())
// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// }

func (u *User) ReadAllContactDetails(contactIdTemp uuid.UUID) (allContactDetailsInfo []*guru_contact_details.ContactDetails) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}

	var requiredContact *guru_contacts.Contact = u.ReadContactById(contactIdTemp)
	if requiredContact == nil {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
	}

	for i := 0; i < len(requiredContact.Contact_Details); i++ {
		flag, contactDetailsTemp := requiredContact.Contact_Details[i].ReadContactDetails()
		if flag {
			allContactDetailsInfo = append(allContactDetailsInfo, contactDetailsTemp)

		}
	}

	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsReadAll).GetSpecificMessage())

}

func (u *User) ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp uuid.UUID) (requiredContactDetails *guru_contact_details.ContactDetails) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}

	var requiredContact *guru_contacts.Contact = u.ReadContactById(contactIdTemp)
	if requiredContact == nil {
		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
	}
	var requiredContactDetailsTemp *guru_contact_details.ContactDetails
	for i := 0; i < len(requiredContact.Contact_Details); i++ {
		if requiredContact.Contact_Details[i].GetContactDetailsId() == contactDetailsIdTemp {
			requiredContactDetailsTemp = requiredContact.Contact_Details[i]
			break
		}
	}
	flag, requiredContactDetailsPrint := requiredContactDetailsTemp.ReadContactDetails()
	if !flag {
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())

	}
	requiredContactDetails = requiredContactDetailsPrint
	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsRead).GetSpecificMessage())

}

func (u *User) DeleteContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {
		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	var requiredContactDetails *guru_contact_details.ContactDetails = u.ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp)
	if requiredContactDetails == nil {
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())
	}

	requiredContactDetails.DeleteContactDetails()
	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeleted).GetSpecificMessage())

}

func (u *User) UpdateContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID, keyName, keyValule string) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	if !u.isActive {

		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

	}
	var requiredContactDetails *guru_contact_details.ContactDetails = u.ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp)
	if requiredContactDetails == nil {
		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())
	}

	requiredContactDetails.UpdateContactDetails(keyName, keyValule)
	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsUpdated).GetSpecificMessage())

}
