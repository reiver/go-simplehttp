package verboten


import (
	"github.com/reiver/go-simplehttp/driver"

	"encoding/json"
	"net/http"
)


const (
	driverName = "json"
)


func init() {
	driver := new(jsonDriver)

	simplehttpdriver.Registry.Register(driverName, driver)
}


type jsonDriver struct{}


func (driver *jsonDriver) Respond(w http.ResponseWriter, httpStatusCode int, httpStatusName string, headers map[string][]string, data map[string]interface{}) {

	// Specify the Content-Type of the HTTP response.
	w.Header().Set("Content-Type", "application/json")

	// Set other headers in the HTTP response.
	for headerName, headerValues := range headers {
		for _, headerValue := range headerValues {
			w.Header().Set(headerName, headerValue)
		}
	}

	// Write the HTTP headers in the response, with the specified HTTP response code.
	w.WriteHeader(httpStatusCode)

	// Add the basic fields.
	//
	// These fields correspond to the HTTP response status code and name.
	data["status_code"] = httpStatusCode
	data["status_name"] = httpStatusName

	// Write out the JSON response.
	jsonEncoder := json.NewEncoder(w)
	if nil == jsonEncoder {
		return
	}

	if err := jsonEncoder.Encode(data); nil != err {
		return
	}
}
