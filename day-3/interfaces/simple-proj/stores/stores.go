package stores

type storer interface {
	Create(usr User) error
	Update(usr User) error
	Delete(usr User) error
}

// StorerInterface is a global variable which will be used by services to interact with store.
// This helps in Dependency Injection and to swap storers without changing code in services.
// this is not recommended because anyone can change the value of the interface
var StorerInterface storer

type Service struct {
	//embedding automatically implement interface if the interface was already implemented by the embedded type
	storer
}

// NewService function is a constructor for Service. It takes an implementer of Storer interface,
// and returns a new Service with the provided Storer. This design allows creating services with
// different storers at runtime.
func NewService(storer storer) Service {
	if storer == nil {
		panic("storer is nil")
	}
	s := Service{storer: storer}
	return s
}
