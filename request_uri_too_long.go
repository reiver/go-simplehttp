package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) RequestURITooLong(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestURITooLong
	httpStatusName :=  StatusNameRequestURITooLong

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
