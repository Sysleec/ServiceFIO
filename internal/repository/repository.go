package repository

import (
	"context"

	"github.com/Sysleec/ServiceFIO/internal/model"
)

// PeopleRepository is the interface that provides people methods.
type PeopleRepository interface {
	Create(ctx context.Context, ppl *model.People) (*model.People, error)
	Get(ctx context.Context, filter string, page int, limit int) ([]*model.People, error)
	Update(ctx context.Context, id int, ppl *model.PeopleReq) (*model.People, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
