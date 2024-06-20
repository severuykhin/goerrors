package goerrors

import "fmt"

type multiError []err

func (me multiError) Error() string {
	message := ""

	for _, er := range me {
		message += fmt.Sprintf("%s; ", er.Error())
	}

	return message
}

func NewMultiError(errors ...error) error {

	if len(errors) == 0 {
		return nil
	}

	innerErrors := make([]err, len(errors))
	for idx, er := range errors {
		innerError := From(er)
		innerErrors[idx] = innerError
	}

	return err{
		kind:        ErrMulti,
		innerErrors: innerErrors,
	}
}
