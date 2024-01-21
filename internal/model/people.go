package model

import (
	"database/sql"
	"time"
)

// People is a model for people
type People struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Surname     string         `json:"surname"`
	Patronymic  sql.NullString `json:"patronymic"`
	Age         int32          `json:"age"`
	Gender      string         `json:"gender"`
	Nationality string         `json:"nationality"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// PeopleReq is a request model for creating a new people
type PeopleReq struct {
	Name       string         `json:"name"`
	Surname    string         `json:"surname"`
	Patronymic sql.NullString `json:"patronymic"`
}
