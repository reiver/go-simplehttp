package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) InternalServerError(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusInternalServerError
	httpStatusName :=  StatusNameInternalServerError

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
