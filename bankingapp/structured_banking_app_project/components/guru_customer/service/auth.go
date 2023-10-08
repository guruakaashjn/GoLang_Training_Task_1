package service

import (
	"bankingapp/models/customer"
	"bankingapp/repository"
)

func (customerService *CustomerService) AuthService(requiredCustomer *customer.Customer, userName string) error {
	uow := repository.NewUnitOfWork(customerService.db, false)
	defer uow.RollBack()
	err := customerService.repository.GetRecord(uow, requiredCustomer, repository.Filter("`user_name` = ?", userName))
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}
