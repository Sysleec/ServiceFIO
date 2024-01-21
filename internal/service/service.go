package service

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// PeopleService is the interface that provides people methods.
type PeopleService interface {
	Create(ctx context.Context, ppl *model.PeopleReq) (*model.People, error)
	Get(ctx context.Context, filter string, page int, limit int) ([]*model.People, error)
	Update(ctx context.Context, id int, ppl *model.PeopleReq) (*model.People, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type EnrichmentService interface {
	GetAge(ctx context.Context, name string) (int, error)
	GetGender(ctx context.Context, name string) (string, error)
	GetNationality(ctx context.Context, name string) (string, error)
}
