package codes

// RecordExists struct is an error that indicates, that record
// you tried to create, already exists
type RecordExists struct{}

func (err *RecordExists) Error() string {
	return "unable to insert value: already exists"
}

// SignalCanceled struct is an error that indicates, that context
// had error value equal to Canceled
type SignalCanceled struct{}

func (err *SignalCanceled) Error() string {
	return "recieved cancel signal"
}

// NotAuthenticated struct is an error that indicates,
// that user cannot be authenticated
type NotAuthenticated struct{}

func (err *NotAuthenticated) Error() string {
	return "username and password mismatch"
}
