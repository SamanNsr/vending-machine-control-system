package errorConstant

var ErrorLists *errorLists

type errorLists struct {
	InternalServerError ErrorList
	ValidationError     ErrorList
	NotFoundError       ErrorList
	BadRequestError     ErrorList
}

type ErrorList struct {
	Msg  string
	Code int
}

func init() {
	ErrorLists = &errorLists{
		// 1000 - 1999 : BoilerPlate Err
		// 2000 - 2999 : Custom Err Per Service
		// .
		// .
		// .
		// 8000 - 8999 : Third-party
		// 9000 - 9999 : FATAL

		InternalServerError: ErrorList{
			Msg:  "internal server error",
			Code: 1000,
		},
		ValidationError: ErrorList{
			Msg:  "request validation failed",
			Code: 1001,
		},
		NotFoundError: ErrorList{
			Msg:  "not found",
			Code: 1002,
		},
		BadRequestError: ErrorList{
			Msg:  "incoming bad request",
			Code: 1003,
		},
	}
}
