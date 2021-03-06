package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) ProxyAuthRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusProxyAuthRequired
	httpStatusName :=  StatusNameProxyAuthRequired

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
