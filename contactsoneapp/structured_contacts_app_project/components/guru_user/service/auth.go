package service

import (
	"contactsoneapp/models/user"
	"contactsoneapp/repository"
)

func (userService *UserService) AuthService(requiredUser *user.User, userName string) error {
	uow := repository.NewUnitOfWork(userService.db, false)
	defer uow.RollBack()
	err := userService.repository.GetRecord(uow, requiredUser, repository.Filter("`user_name` = ?", userName))
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}
