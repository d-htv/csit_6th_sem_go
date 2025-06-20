package c_errors

type CustomError struct{
	Err error
	Message string
	Type string
}

func (ce CustomError) Error() string{
	return "Message : " + ce.Message + " Type : " + ce.Type
}

func Test() error{
	return CustomError{
		Err: nil,
		Message: "test",
		Type: "test",
	}
}