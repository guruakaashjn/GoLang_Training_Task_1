package guru_errors

import "errors"

type invalidUserError struct {
	Error           error
	specificMessage string
}

const NotAnAdminError string = "user is not an admin"
const AdminDeleted string = "admin is deleted"
const UserDeleted string = "user is deleted"

func NewInvalidUserError(specificMessage string) *invalidUserError {
	return &invalidUserError{
		Error:           errors.New("invalid user"),
		specificMessage: specificMessage,
	}

	// return newObjectOfInvalidUser
}

func (e *invalidUserError) GetSpecificMessage() string {
	return e.specificMessage
}

type contactError struct {
	Error           error
	specificMessage string
}

const ContactDeleted string = "contact deleted"
const ContactUpdated string = "contact updated"
const ContactCreated string = "contact created"
const ContactReadAll string = "all contacts read done"
const ContactRead string = "contact read done"

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

const ContactDetailsDeleted string = "contact details deleted"
const ContactDetailsUpdated string = "contact details updated"
const ContactDetailsCreated string = "contact details created"
const ContactDetailsReadAll string = "all contacts details read done"
const ContactDetailsRead string = "contact details read done"

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
