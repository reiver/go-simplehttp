package simplehttpdriver


var (
	Registry Registrar = newRegistrar()
)


type Registrar interface {
	Register(string, Responder)
	Get(string) (Responder, error)
}


type internalRegistrar map[string]Responder


func newRegistrar() Registrar {
	var registrar internalRegistrar = make(map[string]Responder)

	return registrar
}


func (registrar internalRegistrar) Register(name string, responder Responder) {

	registrar[name] = responder
}


func (registrar internalRegistrar) Get(name string) (Responder, error) {

	responder, ok := registrar[name]
	if !ok {
		return nil, errNotFoundComplainer
	}

	return responder, nil
}
