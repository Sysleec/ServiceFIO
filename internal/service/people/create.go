package people

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
	"github.com/Sysleec/ServiceFIO/internal/service/enrich"
)

// Create creates a new people
func (s *serv) Create(ctx context.Context, ppl *model.PeopleReq) (*model.People, error) {
	age, gender, nation, err := enrich.Enrichment(ctx, ppl.Name)
	if err != nil {
		return nil, err
	}

	enrichModel := model.People{
		Name:        ppl.Name,
		Surname:     ppl.Surname,
		Patronymic:  ppl.Patronymic,
		Age:         int32(age),
		Gender:      gender,
		Nationality: nation,
	}
	person, err := s.pplRepo.Create(ctx, &enrichModel)

	if err != nil {
		return nil, err
	}

	return person, nil
}
