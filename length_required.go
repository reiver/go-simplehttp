package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) LengthRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusLengthRequired
	httpStatusName :=  StatusNameLengthRequired

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
