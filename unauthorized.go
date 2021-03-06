package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Unauthorized(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusUnauthorized
	httpStatusName :=  StatusNameUnauthorized

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
