package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) NotFound(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotFound
	httpStatusName :=  StatusNameNotFound

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
