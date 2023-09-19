package errors

type DrawError struct {
	drawEnded string
}

func NewDrawError(specificMessage string) *DrawError {
	return &DrawError{
		drawEnded: specificMessage,
	}
}
func (e *DrawError) GetSpecificMessage() string {
	return e.drawEnded
}
