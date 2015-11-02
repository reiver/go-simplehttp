package simplehttp


import (
	"testing"

	_ "github.com/reiver/go-simplehttp/driver/json"

	"encoding/json"
	"fmt"
	"net/http"
)


// doTest does all the hard work for all the tests.
//
// The actual tests call this.
func doTest(t *testing.T, httpStatusCode int, httpStatusName string, fn func(Responder, http.ResponseWriter, map[string]string)) {

	testName := fmt.Sprintf("%d %s", httpStatusCode, httpStatusName)

	tests := []struct{
		Fields map[string]string
	}{
		{
			Fields: map[string]string{
			},
		},
		{
			Fields: map[string]string{
				"one":"apple",
			},
		},
		{
			Fields: map[string]string{
				"one":"apple",
				"two":"banana",
			},
		},
		{
			Fields: map[string]string{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
			},
		},
	}


	for testNumber, test := range tests {

		const driverName = "json"
		responder, err := Load(driverName)
		if nil != err {
			t.Errorf("For test #%d for %q, did not expect loading of %q driver to return an error but did: %v.", testNumber, testName, driverName, err)
			continue
		}

		mock := newMockResponseWriter()

		// Response specific.
		fn(responder, mock, test.Fields)

		// Response specific.
		if expected, actual := httpStatusCode, mock.StatusCode; expected != actual {
			t.Errorf("For test #%d for %q, expected status code to be %d, but was actually %d.", testNumber, testName, expected, actual)
			continue
		}

		// Make sure response body has some of the stuff we expect in there.
		//
		// For example:
		//	* it is sent as JSON,
		//	* it has a "status_code" field, and
		//	* it has a "status_name" field.
		//
		datum := struct{
			StatusCode int    `json:"status_code"`
			StatusName string `json:"status_name"`
		}{}

		responseBody := mock.Buffer.Bytes();

		if err := json.Unmarshal(responseBody, &datum); nil != err {
			t.Errorf("For test #%d for %q, received an error when trying to unmarshal the response body as JSON: %v", testNumber, testName, err)
			continue
		}

		// Response specific.
		if expected, actual := httpStatusCode, datum.StatusCode; expected != actual {
			t.Errorf("For test #%d for %q, expected status code to be %d, but actually was %d.", testNumber, testName, expected, actual)
			continue
		}

		// Response specific.
		if expected, actual := httpStatusName, datum.StatusName; expected != actual {
			t.Errorf("For test #%d for %q, expected status name to be %q, but actually was %q.", testNumber, testName, expected, actual)
			continue
		}


		// Make sure the reponse body has the extra fields we expect in there.
		softDatum := map[string]interface{}{}
		if err := json.Unmarshal(responseBody, &softDatum); nil != err {
			t.Errorf("For test #%d for %q, received an error when trying to unmarshal the response body as JSON into map[string]interface{}: %v", testNumber, testName, err)
			continue
		}

		if expected, actual := 2 + len(test.Fields), len(softDatum); expected != actual {
			t.Errorf("For test #%d for %q, expected %d fields but actually was %d.", testNumber, testName, expected, actual)
			continue
		}

		for fieldName, fieldValue := range test.Fields {
			if _, ok := softDatum[fieldName]; !ok {
				t.Errorf("For test #%d for %q, expected field name %q to exist, but didn't.", testNumber, testName, fieldName)
				continue
			}

			expected := fieldValue
			if actual, _ := softDatum[fieldName]; expected != actual {
				t.Errorf("For test #%d for %q, for field name %q expected field value to be %q, but was actually %q.", testNumber, testName, fieldName, expected, actual)
				continue
			}
		}
	}
}



func TestOK(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.OK(w, data)
	}

	doTest(t, http.StatusOK, StatusNameOK, fn)
}

func TestAccepted(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Accepted(w, data)
	}

	doTest(t, http.StatusAccepted, StatusNameAccepted, fn)
}

func TestBadGateway(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.BadGateway(w, data)
	}

	doTest(t, http.StatusBadGateway, StatusNameBadGateway, fn)
}

func TestBadRequest(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.BadRequest(w, data)
	}

	doTest(t, http.StatusBadRequest, StatusNameBadRequest, fn)
}

func TestConflict(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Conflict(w, data)
	}

	doTest(t, http.StatusConflict, StatusNameConflict, fn)
}

func TestExpectationFailed(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.ExpectationFailed(w, data)
	}

	doTest(t, http.StatusExpectationFailed, StatusNameExpectationFailed, fn)
}

func TestForbidden(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Forbidden(w, data)
	}

	doTest(t, http.StatusForbidden, StatusNameForbidden, fn)
}

func TestGatewayTimeout(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.GatewayTimeout(w, data)
	}

	doTest(t, http.StatusGatewayTimeout, StatusNameGatewayTimeout, fn)
}

func TestGone(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Gone(w, data)
	}

	doTest(t, http.StatusGone, StatusNameGone, fn)
}

func TestHTTPVersionNotSupported(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.HTTPVersionNotSupported(w, data)
	}

	doTest(t, http.StatusHTTPVersionNotSupported, StatusNameHTTPVersionNotSupported, fn)
}

func TestInternalServerError(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.InternalServerError(w, data)
	}

	doTest(t, http.StatusInternalServerError, StatusNameInternalServerError, fn)
}

func TestLengthRequired(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.LengthRequired(w, data)
	}

	doTest(t, http.StatusLengthRequired, StatusNameLengthRequired, fn)
}

func TestMethodNotAllowed(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.MethodNotAllowed(w, data)
	}

	doTest(t, http.StatusMethodNotAllowed, StatusNameMethodNotAllowed, fn)
}

func TestNotAcceptable(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.NotAcceptable(w, data)
	}

	doTest(t, http.StatusNotAcceptable, StatusNameNotAcceptable, fn)
}

func TestNotFound(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.NotFound(w, data)
	}

	doTest(t, http.StatusNotFound, StatusNameNotFound, fn)
}

func TestNotImplemented(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.NotImplemented(w, data)
	}

	doTest(t, http.StatusNotImplemented, StatusNameNotImplemented, fn)
}

func TestPaymentRequired(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.PaymentRequired(w, data)
	}

	doTest(t, http.StatusPaymentRequired, StatusNamePaymentRequired, fn)
}

func TestPreconditionFailed(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.PreconditionFailed(w, data)
	}

	doTest(t, http.StatusPreconditionFailed, StatusNamePreconditionFailed, fn)
}

func TestProxyAuthRequired(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.ProxyAuthRequired(w, data)
	}

	doTest(t, http.StatusProxyAuthRequired, StatusNameProxyAuthRequired, fn)
}

func TestRequestedRangeNotSatisfiable(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.RequestedRangeNotSatisfiable(w, data)
	}

	doTest(t, http.StatusRequestedRangeNotSatisfiable, StatusNameRequestedRangeNotSatisfiable, fn)
}

func TestRequestEntityTooLarge(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.RequestEntityTooLarge(w, data)
	}

	doTest(t, http.StatusRequestEntityTooLarge, StatusNameRequestEntityTooLarge, fn)
}

func TestRequestTimeout(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.RequestTimeout(w, data)
	}

	doTest(t, http.StatusRequestTimeout, StatusNameRequestTimeout, fn)
}

func TestRequestURITooLong(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.RequestURITooLong(w, data)
	}

	doTest(t, http.StatusRequestURITooLong, StatusNameRequestURITooLong, fn)
}

func TestServiceUnavailable(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.ServiceUnavailable(w, data)
	}

	doTest(t, http.StatusServiceUnavailable, StatusNameServiceUnavailable, fn)
}

func TestTeapot(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Teapot(w, data)
	}

	doTest(t, http.StatusTeapot, StatusNameTeapot, fn)
}

func TestUnauthorized(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.Unauthorized(w, data)
	}

	doTest(t, http.StatusUnauthorized, StatusNameUnauthorized, fn)

}

func TestUnsupportedMediaType(t *testing.T) {
	fn := func(responder Responder, w http.ResponseWriter, data map[string]string) {
		responder.UnsupportedMediaType(w, data)
	}

	doTest(t, http.StatusUnsupportedMediaType, StatusNameUnsupportedMediaType, fn)
}
