-- name: CreatePeople :one
INSERT INTO people (name, surname, patronymic, age, gender, nationality)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
