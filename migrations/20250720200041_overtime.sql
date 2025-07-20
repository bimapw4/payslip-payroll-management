-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS overtime (
    id          UUID        NOT NULL PRIMARY KEY,
    user_id     UUID        ,
    start_time  TIMESTAMPTZ NOT NULL,
    end_time    TIMESTAMPTZ NOT NULL,
    duration    FLOAT       NOT NULL,
    payroll_id  VARCHAR     ,
    created_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by  VARCHAR     NOT NULL,
    updated_by  VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS overtime;
-- +goose StatementEnd
