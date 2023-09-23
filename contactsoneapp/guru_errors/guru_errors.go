package guru_errors

import "errors"

type userError struct {
	Error           error
	specificMessage string
}

const NotAnAdminError string = "user not an admin"
const AdminDeleted string = "admin is deleted"
const AdminCreated string = "admin created successfully"
const UserDeleted string = "user deleted successfully"
const UserDeletedStatus string = "user is deleted"
const UserCreated string = "user created successfully"
const UserUpdated string = "user updated successfully"
const UserRead string = "user read done successfully"
const UserReadAll string = "all user read done successfully"

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

const ContactDeleted string = "contact deleted successfully"
const ContactDeletedStatus string = "contact is deleted"
const ContactUpdated string = "contact updated successfully"
const ContactCreated string = "contact created successfully"
const ContactReadAll string = "all contacts read done successfully"
const ContactRead string = "contact read done successfully"

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

const ContactDetailsDeleted string = "contact details deleted successfully"
const ContactDetailsDeletedStatus string = "contact details is deleted"
const ContactDetailsUpdated string = "contact details updated successfully"
const ContactDetailsCreated string = "contact details created successfully"
const ContactDetailsReadAll string = "all contacts details read done successfully"
const ContactDetailsRead string = "contact details read done successfully"

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
