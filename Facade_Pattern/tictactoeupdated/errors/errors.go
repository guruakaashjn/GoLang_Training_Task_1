package errors

import "errors"

type invalidMove struct {
	errorName       error
	specificMessage string
}

func NewInvalidMove(specificMessage string) *invalidMove {
	return &invalidMove{
		errorName:       errors.New("invalid move"),
		specificMessage: specificMessage,
	}

}

func (e *invalidMove) GetSpecificMessage() string {
	return e.specificMessage
}

type nextTurn struct {
	errorName string
}

func NewNextTurn() *nextTurn {
	return &nextTurn{
		errorName: "next turn",
	}

}
func (e *nextTurn) GetNewNextTurnError() string {
	return e.errorName

}
