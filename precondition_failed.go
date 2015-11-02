package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) PreconditionFailed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusPreconditionFailed
	httpStatusName :=  StatusNamePreconditionFailed

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
