package guru_errors

import "errors"

type invalidUserError struct {
	Error           error
	SpecificMessage string
}

func NewInvalidUserError(SpecificMessage string) *invalidUserError {
	return &invalidUserError{
		Error:           errors.New("Invalid User"),
		SpecificMessage: SpecificMessage,
	}

	// return newObjectOfInvalidUser
}
