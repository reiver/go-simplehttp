package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) NotImplemented(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotImplemented
	httpStatusName :=  StatusNameNotImplemented

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
