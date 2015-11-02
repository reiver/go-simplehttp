/*
Package simplehttp provides a simple interface for sending HTTP responses from within an HTTP handler.

Basic Example

	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		// ...
	
		shttp, err := simplehttp.Load("json")
	
		// ...
	
		// If the input the client send us is not valid,
		// then send a "400 Bad Request".
		if !clientInputValid {
			shttp.BadRequest(w)
			return
		}
	
		// ...
	
		// If we got an error in this case,
		// then send a "500 Internal Server Error".
		if err != nil {
			shttp.InternalServerError(w)
			return
		}
	
		// ...
	
		// Tell the client their request was successful,
		// by sending a "200 OK".
		shttp.OK(w)
	}

Note that in that example, 3 different HTTP responses were shown:
"400 Bad Request", "500 Internal Server Error" and "200 OK".

And perhaps more importantly, each of those was just one line of code! :-)


HTTP Response Customization

By default, simplehttp will send something sensible to the HTTP client as content in the HTTP response.

However, sometimes we may want to include more data in the content of the HTTP response.

For example, a very simple "200 OK" would be....

Code:

	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		// ...
	
		shttp, err := simplehttp.Load("json")
	
		// ...
	
		shttp.OK(w)
	}

Response:

	{
	    "status_code": 200,
	    "status_name": "OK"
	}

The response is simple and sensible. And for may cases it may be sufficient.
But in some instances, we may want to include more data.

There are a number of different ways to include more data. One way is this:

Code:

	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		// ...
	
		shttp, err := simplehttp.Load("json")
	
		// ...
	
		shttp.OK(w, map[string]interface{}{
			"name": "Joe Blow",
			"age":  32,
		})
	}

Response:

	{
	    "status_code": 200,
	    "status_name": "OK",

	    "name": "Joe Blow",
	    "age":  32,
	}

(Note the "name" and "age" fields.)

That previous example used a map. Another way to accomplish the same thing is using a struct, as in:

Code:

	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		// ...
	
		shttp, err := simplehttp.Load("json")

		// ...
	
		shttp.OK(w, struct{
			Name string `key:"name"`
			Age  uint   `key:"age"`
		}{
			Name: "Joe Blow",
			Age   32,
		})
	}

Response:

	{
	    "status_code": 200,
	    "status_name": "OK",

	    "name": "Joe Blow",
	    "age":  32,
	}


JSON, XML, HTML, etc

The previous examples all used JSON. Other formats are available.

(And 3rd parties can even create their own custom formats or replace
the existing ones.)

For the next set of examples, we will switch to sending a "404 Not Found"
HTTP response back to the HTTP client from our ServeHTTP method.

So first, here is a JSON response:

	import (
		_ "github.com/reiver/go-simplehttp/driver/json" // ← JSON
	
		"github.com/reiver/go-simplehttp"
	)
	
	// ...
	
	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		shttp, err := simplehttp.Load("json") // ← JSON
	
		// ...
	
		shttp.NotFound(w)
	}

That JSON content in the response would look like:

	{
	    "status_code": 404,
	    "status_name": "Not Found"
	}

Instead of JSON we could also alternatively send XML content with code like the following:

	import (
		_ "github.com/reiver/go-simplehttp/driver/xml" // ← XML
	
		"github.com/reiver/go-simplehttp"
	)
	
	// ...
	
	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		shttp, err := simplehttp.Load("xml") // ← XML
	
		// ...
	
		shttp.NotFound(w)
	}

That XML content in the response would look like:

	<?xml version="1.0" ?>
	
	<response>
	    <status_code>404</status_code>
	    <status_name>Not Found</status_name>
	</response>

And, instead of XML and JSON, we could send simple HTML with code like the following:

	import (
		_ "github.com/reiver/go-simplehttp/driver/html" // ← HTML
	
		"github.com/reiver/go-simplehttp"
	)
	
	// ...
	
	func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
		shttp, err := simplehttp.Load("html") // ← HTML
	
		// ...
	
		shttp.OK(w)
	}

That HTML content in the response would look like:

	<html>
	  <head>
	    <title>404 Not Found</title>
	  </head>
	  <body>
	    <h1>Not Found</h1>
	    <p>The requested URL was not found.</p>
	  </body>
	</html>

This HTML may not be sufficient for our needs. In the case where we want a lot of control
of the content send to the HTTP client, we could create a new driver.


Custom Drivers

For many needs, such for a HTTP-based (or HTTPS-based) backend API, the builtin
drivers (that output JSON, XML, HTML, etc) are enough. But sometimes we will
want more control of the output.

In those case we will create a custom driver.

A driver implements the simplehttpdriver.Responder interface.

(Note, this should not be confused with the simplehttp.Responder interface,
which is different.)

Here is an example JSON driver that instead of using snake_case for the default
JSON field names, uses camelCase:

	import (
		"github.com/reiver/go-simplehttp/driver"
	
		"encoding/json"
		"net/http"
	)
	
	
	var (
		CamelCaseJsonDriver = new(camelCaseJsonDriver)
	)
	
	
	type camelCaseJsonDriver struct{}
	
	
	func (driver *camelCaseJsonDriver) Respond(w http.ResponseWriter, httpStatusCode int, httpStatusName string, data map[string]interface{}) {
	
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
		data["statusCode"] = httpStatusCode
		data["statusName"] = httpStatusName
	
		// Write out the JSON response.
		jsonEncoder := json.NewEncoder(w)
		if nil == jsonEncoder {
			return
		}
	
		if err := jsonEncoder.Encode(data); nil != err {
			return
		}
	}

To made use of with, we would use code like the following:

	shttp, err := simplehttp.New(CamelCaseJsonDriver)
	
	// ...
	
	shttp.OK(w)

(Note that we called the simplehttp.New func, instead of the simplehttp.Load func.)


Registering Customer Drivers

Although for the times we create a custom driver, many of those times we would
not want to do the following, sometimes you we may want to register our custom
driver so we can make it available to the simplehttp.Load func.

To register a custom driver we would do something like the following:


	import (
		"github.com/reiver/go-simplehttp/driver"
	)
	
	
	const (
		driverName = "camel_case_json"
	)
	
	
	func init() {
		driver := new(camelCaseJsonDriver)
		
		simplehttpdriver.Registry.Register(driverName, driver)
	}
	

It would then be available to load like with the following:

	shttp, err := simplehttp.Load("camel_case_json")
*/
package simplehttp
