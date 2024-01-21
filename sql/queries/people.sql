-- name: CreatePeople :one
INSERT INTO people (name, surname, patronymic, age, gender, nationality)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: DeletePeople :exec
DELETE FROM people WHERE id = $1;


-- name: GetPeople :many
SELECT * FROM people;

-- name: UpdatePeople :one
UPDATE people
SET name = COALESCE($2, name),
    surname = COALESCE($3, surname),
    patronymic = COALESCE($4, patronymic)
WHERE id = $1
RETURNING *;