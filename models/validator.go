package models

// ValidatorError represents a model validation error.
type ValidatorError struct {
	text string
}

// NewValidatorError creates new ValidatorError instance.
func NewValidatorError(text string) *ValidatorError {
	return &ValidatorError{
		text: text,
	}
}

func (e *ValidatorError) Error() string {
	return e.text
}

// IsValidator checks if err is a validator error type.
func IsValidator(err error) bool {
	switch err.(type) {
	case *ValidatorError:
		return true
	default:
		return false
	}
}
