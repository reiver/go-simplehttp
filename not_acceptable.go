package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) NotAcceptable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotAcceptable
	httpStatusName :=  StatusNameNotAcceptable

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
