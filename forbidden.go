package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Forbidden(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusForbidden
	httpStatusName :=  StatusNameForbidden

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
