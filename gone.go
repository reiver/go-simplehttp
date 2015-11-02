package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Gone(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusGone
	httpStatusName :=  StatusNameGone

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
