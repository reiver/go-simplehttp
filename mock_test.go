package simplehttp


import (
	"bytes"
	"net/http"
)



// mockResponseWriter fits the http.ResponseWriter interface, and is used
// in the tests as a 'mock'.
type mockResponseWriter struct {
	Buffer     bytes.Buffer
	StatusCode int
	Headers    http.Header
}

// newMockResponseWriter creates a new mockResponseWriter.
func newMockResponseWriter() *mockResponseWriter {
	headers := make(map[string][]string)

	mock := mockResponseWriter{
		Headers: http.Header(headers),
	}

	return &mock
}

func (mock *mockResponseWriter) Write(b []byte) (int, error) {
	return mock.Buffer.Write(b)
}


func (mock *mockResponseWriter) WriteHeader(statusCode int) {
	mock.StatusCode = statusCode
}

func (mock *mockResponseWriter) Header() http.Header {

	return mock.Headers
}
