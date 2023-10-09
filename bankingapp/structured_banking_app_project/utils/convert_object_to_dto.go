package utils

import "bankingapp/models/customer"

func ConvertUserObjectToDTOObject(customerObj *customer.Customer, customerDTOObj *customer.CustomerDTO) {
	customerDTOObj.FirstName = customerObj.FirstName
	customerDTOObj.LastName = customerObj.LastName
	customerDTOObj.UserName = customerObj.UserName
	customerDTOObj.IsActive = customerObj.IsActive
	customerDTOObj.Accounts = customerObj.Accounts
}
