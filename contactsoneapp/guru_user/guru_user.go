package guru_user

import (
	"contactsoneapp/guru_contacts"
	"fmt"

	"github.com/google/uuid"
)

var userId int = 1

type User struct {
	userId   uuid.UUID
	f_name   string
	l_name   string
	isAdmin  bool
	isActive bool
	Contacts []*guru_contacts.Contact
}

var Users = make([]*User, 0)

func NewUser(f_name, l_name string, isAdmin, isActive bool, contactType string, contactValue string) *User {

	var ContactsTempList = make([]*guru_contacts.Contact, 0)
	var ContactsTempItem = guru_contacts.NewContact(f_name, l_name, isActive, contactType, contactValue)
	var newObjectOfUser = &User{
		userId:   uuid.New(),
		f_name:   f_name,
		l_name:   l_name,
		isAdmin:  isAdmin,
		isActive: isActive,
		Contacts: append(ContactsTempList, ContactsTempItem),
	}

	// fmt.Println(Users)

	return newObjectOfUser
}

// CRUD OPERATIONS

func (u *User) ReadUser() {
	fmt.Println("User First Name: ", u.f_name)
	fmt.Println("User Last Name: ", u.l_name)
	fmt.Println("isAdmin: ", u.isAdmin)
	fmt.Println("isActive: ", u.isActive)
	for i := 0; i < len(u.Contacts); i++ {
		u.Contacts[i].ReadContact()
	}
}

func ReadAllUsers() {
	for i := 0; i < len(Users); i++ {
		Users[i].ReadUser()
	}
	// fmt.Println("")
}

func (u *User) DeleteUser() {
	for i := 0; i < len(u.Contacts); i++ {
		if u.Contacts[i].GetFirstName() == u.f_name {
			u.Contacts[i].DeleteContact()
			break
		}
	}

	u.isActive = false
}

func (u *User) UpdateUser(updateField, updateValue string) {
	// for i := 0; i < len(Users); i++ {
	// 	if Users[i].f_name == firstName {
	// 		switch updateField {
	// 		case "f_name":
	// 			u.f_name = updateValue
	// 		case "l_name":
	// 			u.l_name = updateValue
	// 		}
	// 	}
	// }

	for i := 0; i < len(u.Contacts); i++ {
		if u.Contacts[i].GetFirstName() == u.f_name {
			u.Contacts[i].UpdateContact(updateField, updateValue)
			break
		}
	}

	switch updateField {
	case "f_name":
		u.f_name = updateValue
	case "l_name":
		u.l_name = updateValue
	}

}

func (u *User) CheckIsAdmin() bool {
	return u.isAdmin
}

func (u *User) GetUser() string {
	return u.f_name
}

func GetUserObject(firstName string) *User {
	var tempObj *User
	for i := 0; i < len(Users); i++ {
		if Users[i].GetUser() == firstName {
			tempObj = Users[i]
		}
	}

	return tempObj
}
