package errorhelpers

// ErrorList is a wrapper around many errors
type ErrorList struct {
	start  string
	errors []string
}

// NewErrorList returns a new ErrorList
func NewErrorList(start string) *ErrorList { _ = "STUB: not implemented"; return nil }

// NewErrorListWithErrors returns a new ErrorList with the given errors.
func NewErrorListWithErrors(start string, errors []error) *ErrorList {
	_ = "STUB: not implemented"
	return nil
}

// AddError adds the passed error to the list of errors if it is not nil
func (e *ErrorList) AddError(err error) { _ = "STUB: not implemented"; return }

// AddErrors adds the non-nil errors in the given slice to the list of errors.
func (e *ErrorList) AddErrors(errs ...error) { _ = "STUB: not implemented"; return }

// AddWrap is a convenient wrapper around `AddError(fmt.Errorf("%s: %w", msg, err))`.
func (e *ErrorList) AddWrap(err error, msg string) { _ = "STUB: not implemented"; return }

// AddWrapf is a convenient wrapper around `AddError(fmt.Errorf(format+": %w", append(args, err)...))`.
func (e *ErrorList) AddWrapf(err error, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// AddString adds a string based error to the list
func (e *ErrorList) AddString(err string) { _ = "STUB: not implemented"; return }

// AddStringf adds a templated string
func (e *ErrorList) AddStringf(t string, args ...interface{}) { _ = "STUB: not implemented"; return }

// AddStrings adds multiple string based errors to the list.
func (e *ErrorList) AddStrings(errs ...string) { _ = "STUB: not implemented"; return }

// ToError returns an error if there were errors added or nil
func (e *ErrorList) ToError() error { _ = "STUB: not implemented"; return nil }

// String converts the list to a string, returning empty if no errors were added.
func (e *ErrorList) String() string { _ = "STUB: not implemented"; return "" }

// ErrorStrings returns all the error strings in this ErrorList as a slice, ignoring the start string.
func (e *ErrorList) ErrorStrings() []string { _ = "STUB: not implemented"; return nil }
