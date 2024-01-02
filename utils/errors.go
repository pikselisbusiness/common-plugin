package utils

type LanguageError struct {
	LanguageString  string
	DetailedMessage string
}

func (err LanguageError) Error() string {
	return err.LanguageString
}
func (err LanguageError) LangugeString() string {
	return err.DetailedMessage
}
func (err LanguageError) Details() string {
	return err.DetailedMessage
}
func NewLanguageError(message, detailedMessage string) error {
	return &LanguageError{
		LanguageString:  message,
		DetailedMessage: detailedMessage,
	}
}

type DetailedError struct {
	Message      string
	ErrorDetails error
}

func (err DetailedError) Error() string {
	return err.Message
}
func (err DetailedError) Details() string {
	return err.ErrorDetails.Error()
}

type VariableError struct {
	Message  string
	Variable string
}

func NewVariableError(message, variable string) error {
	return &VariableError{
		Message:  message,
		Variable: variable,
	}
}
func (err VariableError) Error() string {
	return err.Message
}
func (err VariableError) GetVariable() string {
	return err.Variable
}
