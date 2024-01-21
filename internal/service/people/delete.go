package people

import "context"

// Delete deletes a ppl
func (s *serv) Delete(ctx context.Context, id int64) (int64, error) {
	person, err := s.pplRepo.Delete(ctx, id)

	if err != nil {
		return 0, err
	}

	return person, nil
}
