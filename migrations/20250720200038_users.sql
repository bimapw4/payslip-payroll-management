-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS users (
    id          UUID        NOT NULL PRIMARY KEY,
    name        VARCHAR     NOT NULL,
    username    VARCHAR     NOT NULL UNIQUE,
    password    TEXT        NOT NULL,
    salary      INT         NOT NULL,
    is_admin    BOOLEAN     NOT NULL    DEFAULT FALSE,
    is_active   BOOLEAN     NOT NULL    DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by  VARCHAR     NOT NULL,
    updated_by  VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
