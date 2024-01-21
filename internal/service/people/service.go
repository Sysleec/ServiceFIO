package people

import (
	"github.com/Sysleec/ServiceFIO/internal/repository"
	"github.com/Sysleec/ServiceFIO/internal/service"
)

var _ service.PeopleService = (*serv)(nil)

type serv struct {
	pplRepo repository.PeopleRepository
}

// NewService returns a new People service
func NewService(peopleRepo repository.PeopleRepository) *serv {
	return &serv{
		pplRepo: peopleRepo,
	}
}
