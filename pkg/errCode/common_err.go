package errCode

var (
	Success        = NewError(0, "")
	ServerError    = NewError(100, "")
	InvalidParams  = NewError(101, "")
	NotFound       = NewError(102, "")
	TooManyRequest = NewError(103, "")
)

var (
	ErrorCreateShortenURL = NewError(200, "")
	ErrorGetURL           = NewError(201, "")
	ErrorUrlCodeExpired   = NewError(202, "")
)
