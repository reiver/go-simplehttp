package simplehttp


import (
	"net/http"
)


func (responder *internalResponder) PaymentRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusPaymentRequired
	httpStatusName :=  StatusNamePaymentRequired

	data := collapse(cascade...)

	responder.driver.Respond(w, httpStatusCode, httpStatusName, responder.headers, data)
}
