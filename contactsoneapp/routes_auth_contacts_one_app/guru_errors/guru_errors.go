package guru_errors

import "errors"

type userError struct {
	Error           error
	specificMessage string
}

// SERVICE ERRORS
const NotAnAdminError string = "user not an admin"
const AdminDeleted string = "admin is deleted"
const AdminCreated string = "admin created successfully"
const UserDeleted string = "user deleted successfully"
const UserDeletedStatus string = "user is deleted"
const UserCreated string = "user created successfully"
const UserUpdated string = "user updated successfully"
const UserRead string = "user read done successfully"
const UserReadAll string = "all user read done successfully"

// CONTROLLER ERRORS
const CreateAdminFailed string = "create admin failed from user controller"
const CreateUserFailed string = "create user failed from user controller"
const CreateAdminSuccess string = "create admin success from user controller"
const CreateUserSuccess string = "create user success from user controller"
const ReadUserFailed string = "read user failed from user controller"
const ReadUserSuccess string = "read user success from user controller"
const UpdateUserFailed string = "update user failed from user controller"
const UpdateUserSuccess string = "update user success from user controller"
const DeleteUserFailed string = "delete user failed from user controller"
const DeleteUserSuccess string = "delete user success from user controller"

func NewUserError(specificMessage string) *userError {
	return &userError{
		Error:           errors.New("invalid user"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *userError) GetSpecificMessage() string {
	return e.specificMessage
}

type contactError struct {
	Error           error
	specificMessage string
}

// SERVICE ERRORS
const ContactDeleted string = "contact deleted successfully"
const ContactDeletedStatus string = "contact is deleted"
const ContactUpdated string = "contact updated successfully"
const ContactCreated string = "contact created successfully"
const ContactReadAll string = "all contacts read done successfully"
const ContactRead string = "contact read done successfully"

// CONTROLLER ERRORS

const CreateContactFailed string = "create contact failed from contact controller"
const CreateContactSuccess string = "create contact success from contact controller"
const ReadContactFailed string = "read contact failed from contact controller"
const ReadContactSuccess string = "read contact success from contact controller"
const UpdateContactFailed string = "update contact failed from contact controller"
const UpdateContactSuccess string = "update contact success from contact controller"
const DeleteContactFailed string = "delete contact failed from contact controller"
const DeleteContactSuccess string = "delete contact success from contact controller"

func NewContactError(specificMessage string) *contactError {
	return &contactError{
		Error:           errors.New("contact error"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *contactError) GetSpecificMessage() string {
	return e.specificMessage
}

type contactDetailsError struct {
	Error           error
	specificMessage string
}

// SERVICE ERRORS
const ContactDetailsDeleted string = "contact details deleted successfully"
const ContactDetailsDeletedStatus string = "contact details is deleted"
const ContactDetailsUpdated string = "contact details updated successfully"
const ContactDetailsCreated string = "contact details created successfully"
const ContactDetailsReadAll string = "all contacts details read done successfully"
const ContactDetailsRead string = "contact details read done successfully"

// CONTROLLER ERRORS
const CreateContactDetailsFailed string = "create contact details failed from contact details controller"
const CreateContactDetailsSuccess string = "create contact details success from contact details controller"
const ReadContactDetailsFailed string = "read contact details failed from contact details controller"
const ReadContactDetailsSuccess string = "read contact details success from contact details controller"
const UpdateContactDetailsFailed string = "update contact details failed from contact details controller"
const UpdateContactDetailsSuccess string = "update contact details success from contact details controller"
const DeleteContactDetailsFailed string = "delete contact details failed from contact details controller"
const DeleteContactDetailsSuccess string = "delete contact details success from contact details controller"

func NewContactDetailsError(specificMessage string) *contactDetailsError {
	return &contactDetailsError{
		Error:           errors.New("contact details error"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *contactDetailsError) GetSpecificMessage() string {
	return e.specificMessage
}

type authenticationError struct {
	Error           error
	specificMessage string
}

const AuthenticationFailed string = "login failed authentication failed"
const AuthenticationSuccess string = "login success authentication success"

func NewAuthenticationError(specificMessage string) *authenticationError {
	return &authenticationError{
		Error:           errors.New("authentication error"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *authenticationError) GetSpecificMessage() string {
	return e.specificMessage
}
