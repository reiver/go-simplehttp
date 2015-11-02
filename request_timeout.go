package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) RequestTimeout(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestTimeout
	httpStatusName :=  StatusNameRequestTimeout

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
