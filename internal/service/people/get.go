package people

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// Get returns all users
func (s *serv) Get(ctx context.Context, filter string, page int, limit int) ([]*model.People, error) {
	user, err := s.pplRepo.Get(ctx, filter, page, limit)

	if err != nil {
		return nil, err
	}

	return user, nil
}
