package simplehttp


import (
	"github.com/reiver/go-simplehttp/driver"

	"errors"
)


type internalResponder struct {
	driverName  string
	driver      simplehttpdriver.Responder
	headers map[string][]string
}


func Load(name string) (Responder, error) {
	driver, err := simplehttpdriver.Registry.Get(name)
	if nil != err {
//@TODO: Use a typed-error
		return nil, errors.New("Not Found")
	}

	responder := New(driver)
	responder.(*internalResponder).driverName = name

	return responder, nil
}


func New(driver simplehttpdriver.Responder) Responder {

	responder := internalResponder{
		driver:driver,
	}

	return &responder
}
