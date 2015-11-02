package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) RequestEntityTooLarge(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestEntityTooLarge
	httpStatusName :=  StatusNameRequestEntityTooLarge

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
