package simplehttp


import (
	"net/http"
)


// Responder is the interface that contains methods that makes it simple to send
// an HTTP response.
//
// For example, a "404 Not Found" corresponds to the NotFound method, a "500 Internal
// Server Error" corresponds to the InternalServerError method, and a "408 Request
// Timeout" corresponds to the RequestTimeout method. Etc.
type Responder interface {
	OK(http.ResponseWriter, ...interface{})       // 200
	Accepted(http.ResponseWriter, ...interface{}) // 202

	BadRequest(http.ResponseWriter, ...interface{})
	Unauthorized(http.ResponseWriter, ...interface{})
	PaymentRequired(http.ResponseWriter, ...interface{})
	Forbidden(http.ResponseWriter, ...interface{})
	NotFound(http.ResponseWriter, ...interface{})
	MethodNotAllowed(http.ResponseWriter, ...interface{})
	NotAcceptable(http.ResponseWriter, ...interface{})
	ProxyAuthRequired(http.ResponseWriter, ...interface{})
	RequestTimeout(http.ResponseWriter, ...interface{})
	Conflict(http.ResponseWriter, ...interface{})
	Gone(http.ResponseWriter, ...interface{})
	LengthRequired(http.ResponseWriter, ...interface{})
	PreconditionFailed(http.ResponseWriter, ...interface{})
	RequestEntityTooLarge(http.ResponseWriter, ...interface{})
	RequestURITooLong(http.ResponseWriter, ...interface{})
	UnsupportedMediaType(http.ResponseWriter, ...interface{})
	RequestedRangeNotSatisfiable(http.ResponseWriter, ...interface{})
	ExpectationFailed(http.ResponseWriter, ...interface{})
	Teapot(http.ResponseWriter, ...interface{})

	InternalServerError(http.ResponseWriter, ...interface{})
	NotImplemented(http.ResponseWriter, ...interface{})
	BadGateway(http.ResponseWriter, ...interface{})
	ServiceUnavailable(http.ResponseWriter, ...interface{})
	GatewayTimeout(http.ResponseWriter, ...interface{})
	HTTPVersionNotSupported(http.ResponseWriter, ...interface{})
}
