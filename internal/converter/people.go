package converter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sysleec/ServiceFIO/internal/model"
	"github.com/Sysleec/ServiceFIO/internal/repository/people/sqlc"
)

func ToPeopleFromReq(r *http.Request) (*model.PeopleReq, error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	type Params struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	}

	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode JSON: %v", err)
	}

	return &model.PeopleReq{
		Name:       params.Name,
		Surname:    params.Surname,
		Patronymic: toNullString(params.Patronymic),
	}, nil
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToSqlcPersonFromModelPeople(m *model.People) *sqlc.Person {
	return &sqlc.Person{
		ID:          m.ID,
		Name:        m.Name,
		Surname:     m.Surname,
		Patronymic:  m.Patronymic,
		Age:         m.Age,
		Gender:      m.Gender,
		Nationality: m.Nationality,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
