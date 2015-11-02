package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) Conflict(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusConflict
	httpStatusName :=  StatusNameConflict

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
