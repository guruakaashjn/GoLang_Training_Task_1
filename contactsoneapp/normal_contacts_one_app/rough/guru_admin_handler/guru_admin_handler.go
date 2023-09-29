package guru_admin_handler

import (
	"contactsoneapp/guru_errors"
	"contactsoneapp/guru_user"
	"fmt"
)

func AdminHandler() (check bool) {
	// check = false
	fmt.Println("Enter F_Name : ")
	var firstName string
	fmt.Scan(&firstName)
	defer func() {
		if a := recover(); a != nil {
			fmt.Println("Recovered", a)
			check = false
		}

	}()
	check = adminCheck(firstName)
	// fmt.Println("Check: ", check)
	return check
}

func adminCheck(firstName string) bool {
	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName && guru_user.Users[i].CheckIsAdmin() {
			return true
		}
		if guru_user.Users[i].GetUser() == firstName {
			msg := guru_user.Users[i].GetUser() + " is not an Admin!"
			panic(guru_errors.NewInvalidUserError(msg))
		}
	}
	msg := firstName + " does not exist!"
	panic(guru_errors.NewInvalidUserError(msg))

}

func AdminPrivilages() {

	for i := 0; i < 1; {
		// fmt.Println("Length: ", len(guru_user.Users))
		fmt.Println("\n\nMenu:")
		fmt.Println("1. Create a User")
		fmt.Println("2. Read All Users")
		fmt.Println("3. Update a User")
		fmt.Println("4. Delete a user")
		fmt.Println("5. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var userObj1 *guru_user.User
			// fmt.Println("1 : ", len(guru_user.Users))
			userObj1 = getInputFromUser()
			newUsers := append(guru_user.Users, userObj1)
			guru_user.Users = newUsers
			// fmt.Println("2 : ", len(guru_user.Users))
		case 2:
			guru_user.ReadAllUsers()
		case 3:
			updateUser()

		case 4:
			deleteUser()
		case 5:
			i++

		}
	}
}

func deleteUser() {
	fmt.Println("Enter Firstname: ")
	var firstName string
	fmt.Scan(&firstName)

	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			guru_user.Users[i].DeleteUser()
		}
	}
}

func updateUser() {
	fmt.Println("Enter Firstname: ")
	var firstName string
	fmt.Scan(&firstName)
	fmt.Println("Enter field to update: ")
	var fieldName string
	fmt.Scan(&fieldName)
	fmt.Println("Enter field value to update: ")
	var fieldValue string
	fmt.Scan(&fieldValue)

	for i := 0; i < len(guru_user.Users); i++ {
		if guru_user.Users[i].GetUser() == firstName {
			guru_user.Users[i].UpdateUser(fieldName, fieldValue)
			break
		}
	}
}

func getInputFromUser() *guru_user.User {
	var firstName string
	var lastName string
	var isAdmin bool
	var isActive bool
	var contactType string
	var contactValue string
	var adminChoice int
	var contactChoice int

	fmt.Println("Enter firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter isAdmin: ")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Scan(&adminChoice)
	switch adminChoice {
	case 1:
		isAdmin = true
	case 2:
		isAdmin = false
	}

	// fmt.Println("Enter isActive: ")
	// fmt.Scan(&isActive)
	isActive = true

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

	userObj1 := guru_user.NewUser(firstName, lastName, isAdmin, isActive, contactType, contactValue)
	return userObj1

}
