package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) ServiceUnavailable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusServiceUnavailable
	httpStatusName :=  StatusNameServiceUnavailable

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
