package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) BadGateway(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusBadGateway
	httpStatusName :=  StatusNameBadGateway

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
