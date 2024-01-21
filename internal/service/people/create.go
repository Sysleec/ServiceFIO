package people

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// Create creates a new people
func (s *serv) Create(ctx context.Context, ppl *model.PeopleReq) (*model.People, error) {
	user, err := s.pplRepo.Create(ctx, ppl)

	if err != nil {
		return nil, err
	}

	//fmt.Println("User created in service")
	return user, nil
}
