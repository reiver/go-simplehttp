package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Teapot(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusTeapot
	httpStatusName :=  StatusNameTeapot

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}