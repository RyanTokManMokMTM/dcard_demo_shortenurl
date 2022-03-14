package errCode

//Shorten URL error Code prefix starting with 1000x
var (
	Success             = NewError(0, "succeed")
	PermanentlyRedirect = NewError(1, "redirect permanently")
	ServerError         = NewError(10000, "server internal error")
	InvalidParams       = NewError(10001, "invalid params")
	NotFound            = NewError(10002, "not found")
	TooManyRequest      = NewError(10003, "too many request")
)

//Shorten URL error Code prefix starting with 2000x
var (
	ErrorCreateShortenURL = NewError(20001, "error create shorten url")
	ErrorGetURL           = NewError(20002, "error get url")
	ErrorUrlCodeExpired   = NewError(20003, "error url Code expired")
)
