package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) GatewayTimeout(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusGatewayTimeout
	httpStatusName :=  StatusNameGatewayTimeout

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
