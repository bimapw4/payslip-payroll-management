-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendance (
    id          UUID        NOT NULL PRIMARY KEY,
    user_id     UUID        ,
    check_in    TIMESTAMPTZ NOT NULL,
    check_out   TIMESTAMPTZ,
    payroll_id  VARCHAR     ,
    created_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by  VARCHAR     NOT NULL,
    updated_by  VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attendance;
-- +goose StatementEnd
