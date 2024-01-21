package converter

import (
	"github.com/Sysleec/ServiceFIO/internal/model"
	modelRepo "github.com/Sysleec/ServiceFIO/internal/repository/people/sqlc"
)

func FromSqlcPersonToModelPeople(person modelRepo.Person) *model.People {
	return &model.People{
		ID:          person.ID,
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Age:         person.Age,
		Gender:      person.Gender,
		Nationality: person.Nationality,
		CreatedAt:   person.CreatedAt,
		UpdatedAt:   person.UpdatedAt,
	}
}
