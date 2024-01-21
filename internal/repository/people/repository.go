package people

import (
	"context"
	"database/sql"

	"github.com/Sysleec/ServiceFIO/internal/model"
	"github.com/Sysleec/ServiceFIO/internal/repository"
	"github.com/Sysleec/ServiceFIO/internal/repository/people/converter"
	"github.com/Sysleec/ServiceFIO/internal/repository/people/sqlc"
)

type repo struct {
	DB *sqlc.Queries
}

// NewRepo creates new people repository
func NewRepo(db *sqlc.Queries) repository.PeopleRepository {
	return &repo{DB: db}
}

// Create creates a new people
func (r *repo) Create(ctx context.Context, ppl *model.PeopleReq) (*model.People, error) {
	user, err := r.DB.CreatePeople(ctx, sqlc.CreatePeopleParams{
		Name:        ppl.Name,
		Surname:     ppl.Surname,
		Patronymic:  ppl.Patronymic,
		Age:         20,
		Gender:      "Male",
		Nationality: "Russian",
	})

	if err != nil {
		return nil, err
	}

	//fmt.Printf("User created in repo: %v\n", user)
	return converter.FromSqlcPersonToModelPeople(user), nil
}

// Get returns all ppls
func (r *repo) Get(ctx context.Context, filter string, page int, limit int) ([]*model.People, error) {

	users, err := r.DB.GetPeople(ctx, sqlc.GetPeopleParams{
		Column1: sql.NullString{String: filter, Valid: true},
		Limit:   int32(limit),
		Column3: int32(page),
	},
	)

	if err != nil {
		return nil, err
	}

	//fmt.Printf("User created in repo: %v\n", user)
	return converter.FromSqlcPersonSliceToModelPeopleSlice(users), nil
}

// Update updates a ppl
func (r *repo) Update(ctx context.Context, ppl *model.PeopleReq) (int64, error) {
	return 0, nil
}

// Delete deletes a ppl
func (r *repo) Delete(ctx context.Context, id int64) (int64, error) {
	err := r.DB.DeletePeople(ctx, int32(id))
	if err != nil {
		return 0, err
	}
	return id, nil
}
