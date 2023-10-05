package service

import (
	"contactsoneapp/models/user"
	"contactsoneapp/repository"
	"contactsoneapp/utils"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

func NewUserService(db *gorm.DB, repo repository.Repository) *UserService {
	return &UserService{
		db:           db,
		repository:   repo,
		associations: []string{"Contacts", "Contacts.ContactInfos"},
	}
}

func (userService *UserService) doesUserExist(Id uint) error {
	exists, err := repository.DoesRecordExistForUser(userService.db, int(Id), user.User{}, repository.Filter("`id` = ?", Id))

	if !exists || err != nil {
		return errors.New("data id is invalid")
	}
	return nil
}
func (userService *UserService) CreateUser(newUser *user.User) error {
	uow := repository.NewUnitOfWork(userService.db, false)
	defer uow.RollBack()

	HashedPassword, _ := utils.GenerateHash(newUser.Password)
	newUser.Password = string(HashedPassword)

	err := userService.repository.Add(uow, newUser)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}
func (userService *UserService) GetAllUsers(allUsers *[]user.User, totalCount *int) error {
	uow := repository.NewUnitOfWork(userService.db, true)
	defer uow.RollBack()
	err := userService.repository.GetAll(uow, allUsers, userService.associations)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
func (userService *UserService) GetUserById(requiredUser *user.User, idTemp int) error {
	uow := repository.NewUnitOfWork(userService.db, true)
	defer uow.RollBack()

	err := userService.repository.GetRecordForId(uow, uint(idTemp), requiredUser)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
func (userService *UserService) UpdateUser(userToUpdate *user.User) error {
	err := userService.doesUserExist(userToUpdate.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(userService.db, false)
	defer uow.RollBack()
	tempUser := user.User{}
	err = userService.repository.GetRecordForId(uow, userToUpdate.ID, &tempUser, repository.Select("`created_at`"), repository.Filter("`id` = ?", userToUpdate.ID))
	if err != nil {
		return err
	}
	userToUpdate.CreatedAt = tempUser.CreatedAt
	err = userService.repository.Save(uow, userToUpdate)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}

func (userService *UserService) DeleteUser(userToDelete *user.User) error {
	err := userService.doesUserExist(userToDelete.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(userService.db, false)
	defer uow.RollBack()

	if err := userService.repository.UpdateWithMap(uow, userToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
	},
		repository.Filter("`id`=?", userToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

// package service

// import (
// 	contact_details_service "contactsoneapp/components/guru_contact_details/service"
// 	"contactsoneapp/utils"

// 	// contacts_details_service "contactsoneapp/components/guru_contact_details/service"
// 	contacts_service "contactsoneapp/components/guru_contacts/service"
// 	"contactsoneapp/guru_errors"
// 	"fmt"

// 	"github.com/google/uuid"
// )

// // var UserId int = 1

// type User struct {
// 	UserId    uuid.UUID
// 	FirstName string
// 	LastName  string
// 	UserName  string
// 	Password  string
// 	IsAdmin   bool
// 	IsActive  bool
// 	Contacts  []*contacts_service.Contact
// }

// var Users = make([]*User, 0)

// func NewUser(FirstName, LastName string, UserName string, Password string, IsAdmin bool) *User {

// 	var ContactsTempList []*contacts_service.Contact = make([]*contacts_service.Contact, 0)
// 	// var ContactsTempItem = guru_contacts.NewContact(FirstName, LastName, true, contactType, contactValue)
// 	HashedPassword, _ := utils.GenerateHash(Password)
// 	var newObjectOfUser = &User{
// 		UserId:    uuid.New(),
// 		FirstName: FirstName,
// 		LastName:  LastName,
// 		UserName:  UserName,
// 		Password:  string(HashedPassword),
// 		IsAdmin:   IsAdmin,
// 		IsActive:  true,
// 		// Contacts: append(ContactsTempList, ContactsTempItem),
// 		Contacts: ContactsTempList,
// 	}
// 	Users = append(Users, newObjectOfUser)

// 	// fmt.Println(Users)

// 	return newObjectOfUser
// }

// // ADMIN CRUD OPERATIONS ON USERS
// // ADMIN
// func CreateAdmin(FirstName, LastName, UserName, Password string) (adminUser *User) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	adminUser = NewUser(FirstName, LastName, UserName, Password, true)
// 	panic(guru_errors.NewUserError(guru_errors.AdminCreated).GetSpecificMessage())

// }

// func (u *User) CreateUser(FirstName, LastName, UserName, Password string) (newUser *User) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// 	}
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	newUser = NewUser(FirstName, LastName, UserName, Password, false)
// 	panic(guru_errors.NewUserError(guru_errors.UserCreated).GetSpecificMessage())

// }

// func (u *User) ReadUserById(UserIdTemp uuid.UUID) (userInfo *User) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

// 	}
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	var requiredUser *User
// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].UserId == UserIdTemp {
// 			requiredUser = Users[i]
// 			break
// 		}
// 	}
// 	if !requiredUser.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

// 	}
// 	userInfo = requiredUser
// 	panic(guru_errors.NewUserError(guru_errors.UserRead).GetSpecificMessage())

// }

// func (u *User) ReadAllUsers() (allUserInfo []*User) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())

// 	}

// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].IsActive {
// 			allUserInfo = append(allUserInfo, Users[i])
// 		}

// 	}
// 	panic(guru_errors.NewUserError(guru_errors.UserReadAll).GetSpecificMessage())

// }

// func (u *User) DeleteUser(UserIdTemp uuid.UUID) (deletedUser *User) {
// 	// UserIdUuid := uuid.Must(uuid.FromBytes([]byte(UserId)))

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// 	}
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	requiredUser := u.ReadUserById(UserIdTemp)
// 	if requiredUser == nil {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

// 	}

// 	for i := 0; i < len(requiredUser.Contacts); i++ {
// 		requiredUser.Contacts[i].DeleteContact()
// 	}

// 	requiredUser.IsActive = false
// 	deletedUser = requiredUser
// 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// }

// func (u *User) UpdateUser(UserIdTemp uuid.UUID, updateField, updateValue string) (updatedUser *User) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// 	}
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	requiredUser := u.ReadUserById(UserIdTemp)
// 	if requiredUser == nil {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())
// 	}

// 	switch updateField {
// 	case "FirstName":
// 		requiredUser.FirstName = updateValue
// 	case "LastName":
// 		requiredUser.LastName = updateValue
// 	}
// 	updatedUser = requiredUser
// 	panic(guru_errors.NewUserError(guru_errors.UserUpdated).GetSpecificMessage())

// }

// func (u *User) UpdateUserObject(UserIdTemp uuid.UUID, userTempObject *User) (updatedUser *User) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !u.IsAdmin {
// 		panic(guru_errors.NewUserError(guru_errors.NotAnAdminError).GetSpecificMessage())
// 	}
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.AdminDeleted).GetSpecificMessage())
// 	}

// 	requiredUser := u.ReadUserById(UserIdTemp)
// 	if requiredUser == nil {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())
// 	}

// 	if userTempObject.FirstName != "" && userTempObject.FirstName != requiredUser.FirstName {
// 		requiredUser.FirstName = userTempObject.FirstName
// 	}
// 	if userTempObject.LastName != "" && userTempObject.LastName != requiredUser.LastName {
// 		requiredUser.LastName = userTempObject.LastName
// 	}
// 	if userTempObject.UserName != "" && userTempObject.UserName != requiredUser.UserName {
// 		requiredUser.UserName = userTempObject.UserName
// 	}
// 	updatedUser = requiredUser
// 	// fmt.Println("Inside Update User Object (updated User): ", updatedUser)
// 	panic(guru_errors.NewUserError(guru_errors.UserUpdated).GetSpecificMessage())

// }

// func (u *User) CheckIsAdmin() bool {
// 	return u.IsAdmin
// }

// func (u *User) GetUser() string {
// 	return u.FirstName
// }

// // USER CRUD OPERATIONS ON CONTACTS
// // USER
// func (u *User) CreateContact(FirstName, LastName string) (newContact *contacts_service.Contact) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	newContact = contacts_service.CreateContact(FirstName, LastName)
// 	u.Contacts = append(u.Contacts, newContact)
// 	panic(guru_errors.NewContactError(guru_errors.ContactCreated).GetSpecificMessage())

// }

// // func (u *User) ReadContactById(contactIdTemp uuid.UUID) (contactInfo string) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()
// // 	if u.IsActive {
// // 		var requiredContact *contacts_service.Contact = u.GetRequiredContactObjectById(contactIdTemp)
// // 		contactInfo = requiredContact.ReadContact()
// // 		panic(guru_errors.NewContactError(guru_errors.ContactRead).GetSpecificMessage())

// // 	}
// // 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// // }
// // func (u *User) ReadAllContact() (allContactInfo string) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()

// // 	if u.IsActive {
// // 		for i := 0; i < len(u.Contacts); i++ {
// // 			allContactInfo += u.Contacts[i].ReadContact() + "\n"
// // 		}
// // 		panic(guru_errors.NewContactError(guru_errors.ContactReadAll).GetSpecificMessage())

// // 	}
// // 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())
// // }

// func (u *User) ReadContactById(contactIdTemp uuid.UUID) (contactInfo *contacts_service.Contact) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContact *contacts_service.Contact

// 	for i := 0; i < len(u.Contacts); i++ {
// 		if u.Contacts[i].GetContactId() == contactIdTemp {
// 			requiredContact = u.Contacts[i]
// 			break
// 		}
// 	}

// 	if !requiredContact.GetIsActive() {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())

// 	}
// 	contactInfo = requiredContact
// 	panic(guru_errors.NewContactError(guru_errors.ContactRead).GetSpecificMessage())

// }
// func (u *User) ReadAllContact() (allContactInfo []*contacts_service.Contact) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	if !u.IsActive {

// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	for i := 0; i < len(u.Contacts); i++ {
// 		flag, contactTemp := u.Contacts[i].ReadContact()
// 		if flag {
// 			allContactInfo = append(allContactInfo, contactTemp)
// 		}

// 	}
// 	panic(guru_errors.NewContactError(guru_errors.ContactReadAll).GetSpecificMessage())

// }

// func (u *User) DeleteContact(contactIdTemp uuid.UUID) (deletedContact *contacts_service.Contact) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// 	if requiredContact == nil {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// 	}

// 	deletedContact = requiredContact.DeleteContact()
// 	panic(guru_errors.NewContactError(guru_errors.ContactDeleted).GetSpecificMessage())

// }

// // func (u *User) UpdateContact(contactIdTemp uuid.UUID, updateField, updateValue string) (updatedContact *contacts_service.Contact) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()
// // 	if !u.IsActive {
// // 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// // 	}

// // 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// // 	if requiredContact == nil {
// // 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// // 	}

// // 	updatedContact = requiredContact.UpdateContact(updateField, updateValue)
// // 	panic(guru_errors.NewUserError(guru_errors.ContactUpdated).GetSpecificMessage())

// // }

// func (u *User) UpdateContactObject(contactIdTemp uuid.UUID, contactTempObj *contacts_service.Contact) (updatedContact *contacts_service.Contact) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}

// 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// 	if requiredContact == nil {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// 	}

// 	updatedContact = requiredContact.UpdateContactObject(contactTempObj)
// 	panic(guru_errors.NewUserError(guru_errors.ContactUpdated).GetSpecificMessage())

// }

// // USER CRUD OPERATIONS ON CONTACT DETAILS
// // USER
// func (u *User) CreateContactDetails(contactIdTemp uuid.UUID, typeName, typeValue string) (newContactDetails *contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// 	if requiredContact == nil {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// 	}

// 	newContactDetails = contact_details_service.CreateContactDetails(typeName, typeValue)
// 	requiredContact.Contact_Details = append(requiredContact.Contact_Details, newContactDetails)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsCreated).GetSpecificMessage())

// }

// // func (u *User) ReadAllContactDetails(contactIdTemp uuid.UUID) (allContactDetailsInfo string) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()
// // 	if u.IsActive {
// // 		var requiredContact *contacts_service.Contact = u.GetRequiredContactObjectById(contactIdTemp)

// // 		allContactDetailsInfo = requiredContact.ReadContact()

// // 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsReadAll).GetSpecificMessage())
// // 	}
// // 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// // }

// // func (u *User) ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp uuid.UUID) (contactDetailsInfo string) {
// // 	defer func() {
// // 		if a := recover(); a != nil {
// // 			fmt.Println(a)
// // 		}
// // 	}()
// // 	if u.IsActive {
// // 		var requiredContactDetails *contact_details_service.ContactDetails = u.GetRequiredContactDetailsObjectById(contactIdTemp, contactDetailsIdTemp)
// // 		contactDetailsInfo = requiredContactDetails.ReadContactDetails()
// // 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsRead).GetSpecificMessage())
// // 	}
// // 	panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// // }

// func (u *User) ReadAllContactDetails(contactIdTemp uuid.UUID) (allContactDetailsInfo []*contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}

// 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// 	if requiredContact == nil {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// 	}

// 	for i := 0; i < len(requiredContact.Contact_Details); i++ {
// 		flag, contactDetailsTemp := requiredContact.Contact_Details[i].ReadContactDetails()
// 		if flag {
// 			allContactDetailsInfo = append(allContactDetailsInfo, contactDetailsTemp)

// 		}
// 	}

// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsReadAll).GetSpecificMessage())

// }

// func (u *User) ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp uuid.UUID) (requiredContactDetails *contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}

// 	var requiredContact *contacts_service.Contact = u.ReadContactById(contactIdTemp)
// 	if requiredContact == nil {
// 		panic(guru_errors.NewContactError(guru_errors.ContactDeletedStatus).GetSpecificMessage())
// 	}
// 	var requiredContactDetailsTemp *contact_details_service.ContactDetails
// 	for i := 0; i < len(requiredContact.Contact_Details); i++ {
// 		if requiredContact.Contact_Details[i].GetContactDetailsId() == contactDetailsIdTemp {
// 			requiredContactDetailsTemp = requiredContact.Contact_Details[i]
// 			break
// 		}
// 	}
// 	flag, requiredContactDetailsPrint := requiredContactDetailsTemp.ReadContactDetails()
// 	if !flag {
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())

// 	}
// 	requiredContactDetails = requiredContactDetailsPrint
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsRead).GetSpecificMessage())

// }

// func (u *User) DeleteContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID) (deletedContactDetails *contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContactDetails *contact_details_service.ContactDetails = u.ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp)
// 	if requiredContactDetails == nil {
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())
// 	}

// 	deletedContactDetails = requiredContactDetails.DeleteContactDetails()
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeleted).GetSpecificMessage())

// }

// func (u *User) UpdateContactDetails(contactIdTemp, contactDetailsIdTemp uuid.UUID, keyName, keyValule string) (updatedContactDetails *contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {

// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContactDetails *contact_details_service.ContactDetails = u.ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp)
// 	if requiredContactDetails == nil {
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())
// 	}

// 	updatedContactDetails = requiredContactDetails.UpdateContactDetails(keyName, keyValule)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsUpdated).GetSpecificMessage())

// }

// func (u *User) UpdateContactDetailsObject(contactIdTemp, contactDetailsIdTemp uuid.UUID, contactDetailsTempObj *contact_details_service.ContactDetails) (updatedContactDetails *contact_details_service.ContactDetails) {
// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()
// 	if !u.IsActive {

// 		panic(guru_errors.NewUserError(guru_errors.UserDeleted).GetSpecificMessage())

// 	}
// 	var requiredContactDetails *contact_details_service.ContactDetails = u.ReadContactDetailsById(contactIdTemp, contactDetailsIdTemp)
// 	if requiredContactDetails == nil {
// 		panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsDeletedStatus).GetSpecificMessage())
// 	}

// 	updatedContactDetails = requiredContactDetails.UpdateContactDetailsObject(contactDetailsTempObj)
// 	panic(guru_errors.NewContactDetailsError(guru_errors.ContactDetailsUpdated).GetSpecificMessage())

// }

// func ReadUserById(userIdTemp uuid.UUID) (requiredUser *User) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}()

// 	var requiredUserTemp *User
// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].UserId == userIdTemp {
// 			requiredUserTemp = Users[i]
// 			break
// 		}
// 	}
// 	if !requiredUserTemp.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

// 	}
// 	requiredUser = requiredUserTemp
// 	panic(guru_errors.NewUserError(guru_errors.UserRead).GetSpecificMessage())
// }

// func ReadUserByUserName(userNameTemp string) (requiredUser *User) {

// 	defer func() {
// 		if a := recover(); a != nil {
// 			fmt.Println(a)
// 		}
// 	}

// 	var requiredUserTemp *User
// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].UserName == userNameTemp {
// 			requiredUserTemp = Users[i]
// 			break
// 		}
// 	}
// 	if !requiredUserTemp.IsActive {
// 		panic(guru_errors.NewUserError(guru_errors.UserDeletedStatus).GetSpecificMessage())

// 	}
// 	requiredUser = requiredUserTemp
// 	panic(guru_errors.NewUserError(guru_errors.UserRead).GetSpecificMessage())
// }
