package people

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// Create creates a new people
func (s *serv) Create(ctx context.Context, ppl *model.PeopleReq) (*model.People, error) {
	person, err := s.pplRepo.Create(ctx, ppl)

	if err != nil {
		return nil, err
	}

	return person, nil
}
