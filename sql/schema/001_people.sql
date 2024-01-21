-- +goose Up
CREATE TABLE IF NOT EXISTS
    people (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        surname VARCHAR(255) NOT NULL,
        patronymic VARCHAR(255),
        age INT NOT NULL DEFAULT 0,
        gender VARCHAR(10) NOT NULL,
        nationality VARCHAR(50) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down
DROP TABLE IF EXISTS people;