package simplehttpdriver


import (
	"net/http"
)


type Responder interface {
	Respond(w http.ResponseWriter, httpStatusCode int, httpStatusName string, headers map[string][]string, data map[string]interface{})
}
