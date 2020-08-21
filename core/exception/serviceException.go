package exception

type ServiceException struct {
	code string
	msg string
	tip string
}

func (e *ServiceException) Error() string {
	return e.msg
}

func NewServiceExceptionWithTip(code string, msg string, tip string) *ServiceException {
	return &ServiceException{
		code: code,
		msg: msg,
		tip: tip,
	}
}


func NewServiceException(code string, msg string) *ServiceException {
	return &ServiceException{
		code: code,
		msg: msg,
		tip: msg,
	}
}
