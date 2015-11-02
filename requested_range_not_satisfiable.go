package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) RequestedRangeNotSatisfiable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestedRangeNotSatisfiable
	httpStatusName :=  StatusNameRequestedRangeNotSatisfiable

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
