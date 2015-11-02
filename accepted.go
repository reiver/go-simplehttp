package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Accepted(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusAccepted
	httpStatusName :=  StatusNameAccepted

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
