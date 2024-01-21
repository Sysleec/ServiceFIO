package people

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// Update updates a ppl
func (s *serv) Update(ctx context.Context, id int, ppl *model.PeopleReq) (*model.People, error) {
	person, err := s.pplRepo.Update(ctx, id, ppl)

	if err != nil {
		return nil, err
	}

	return person, nil
}
