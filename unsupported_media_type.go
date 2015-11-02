package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) UnsupportedMediaType(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusUnsupportedMediaType
	httpStatusName :=  StatusNameUnsupportedMediaType

	data := collapse(responder.driverName, cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
