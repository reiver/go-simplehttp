package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) ExpectationFailed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusExpectationFailed
	httpStatusName :=  StatusNameExpectationFailed

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
