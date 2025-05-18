-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    header TEXT NOT NULL,
    experience TEXT NOT NULL,
    employment TEXT NOT NULL,
    schedule TEXT NOT NULL,
    work_format TEXT NOT NULL,
    working_hours TEXT NOT NULL,
    description TEXT NOT NULL,
    employer_id INT NOT NULL,
    CONSTRAINT fk_employer FOREIGN KEY (employer_id) REFERENCES employers(id)
);

CREATE TABLE IF NOT EXISTS workers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT
);

CREATE TABLE IF NOT EXISTS resume (
    id SERIAL PRIMARY KEY,
    worker_id INT NOT NULL,
    about TEXT NOT NULL,
    experience TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
