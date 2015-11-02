package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) BadRequest(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusBadRequest
	httpStatusName :=  StatusNameBadRequest

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
