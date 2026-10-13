-- +goose Up
CREATE TABLE users (
    id           TEXT        PRIMARY KEY, -- sub de Cognito
    display_name TEXT        NOT NULL DEFAULT '',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users;
