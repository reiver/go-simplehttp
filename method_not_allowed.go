package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) MethodNotAllowed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusMethodNotAllowed
	httpStatusName :=  StatusNameMethodNotAllowed

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
