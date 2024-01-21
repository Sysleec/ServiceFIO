package people

import (
	"github.com/Sysleec/ServiceFIO/internal/service"
)

// Implementation is the people service
type Implementation struct {
	peopleService service.PeopleService
}

// NewImplementation returns a new people service
func NewImplementation(pplService service.PeopleService) *Implementation {
	return &Implementation{peopleService: pplService}
}
