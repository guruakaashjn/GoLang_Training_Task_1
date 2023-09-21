package guru_handler

import (
	"contactsoneapp/guru_admin_handler"
	"contactsoneapp/guru_user_handler"
	"fmt"
)

func MainHandler() {
	for i := 0; i < 1; {

		fmt.Println("Menu:")
		fmt.Println("1. Admin Login")
		fmt.Println("2. User Login")
		fmt.Println("3. Exit")
		var loginChoice int
		fmt.Scan(&loginChoice)
		switch loginChoice {
		case 1:
			// fmt.Println("Case 1: ", guru_admin_handler.AdminHandler())
			if guru_admin_handler.AdminHandler() {
				guru_admin_handler.AdminPrivilages()
				// fmt.Println("Called: ")
			}
		case 2:
			flag, firstName := guru_user_handler.UserHandler()
			if flag {
				guru_user_handler.UserPrivilages(firstName)
			}
		case 3:
			i++
		case 4:
			// For test cases only
			guru_admin_handler.AdminPrivilages()
		}
	}
}
