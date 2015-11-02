package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) HTTPVersionNotSupported(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusHTTPVersionNotSupported
	httpStatusName :=  StatusNameHTTPVersionNotSupported

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
