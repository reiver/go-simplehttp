# go-simplehttp

A library that provides a simple way of sending an HTTP (and HTTPS) response, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-simplehttp

[![GoDoc](https://godoc.org/github.com/reiver/go-simplehttp?status.svg)](https://godoc.org/github.com/reiver/go-simplehttp)



## Example

Sending a "200 OK" HTTP response.
```
simplehttp.OK(w)
```


Sending a "200 OK" HTTP response, with some data.
```
simplehttp.OK(w, struct{
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
}{
    FullName: "Joe Blow",
    Age:      34,
})
```


Sending a "500 Internal Server Error" HTTP response.
```
simplehttp.InternalServerError(w)
```


Sending a "404 Not Found" HTTP response.
```
simplehttp.NotFound(w, map[string]interface{}{
    "message":   "What you are looking for is not here.", 
    "date_time": time.Now(),
})
```
