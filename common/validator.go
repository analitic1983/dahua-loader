package common

type ValidatorError struct {
	Field   string
	Message string
}

type ValidatorErrors struct {
	Errors []ValidatorError
}

type Stringer interface {
	String() string
}

func (v ValidatorErrors) Error() string {
	resultString := "Validation errors: \n"
	for _, oneError := range v.Errors {
		resultString += oneError.Field + ": " + oneError.Message + "\n"
	}
	return resultString
}

func (v *ValidatorErrors) Add(fieldName string, message string) {
	validatorError := ValidatorError{fieldName, message}
	v.Errors = append(v.Errors, validatorError)
}

func (v ValidatorErrors) HasErrors() bool {
	return len(v.Errors) > 0
}
