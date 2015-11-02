package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) OK(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusOK
	httpStatusName :=  StatusNameOK

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
