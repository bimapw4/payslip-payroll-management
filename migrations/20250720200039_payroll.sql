-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payroll (
    id              UUID        NOT NULL    PRIMARY KEY,
    period_start    TIMESTAMPTZ NOT NULL,
    period_end      TIMESTAMPTZ NOT NULL,
    run_payroll     BOOLEAN     NOT NULL    DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by      VARCHAR     NOT NULL,
    updated_by      VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payroll;
-- +goose StatementEnd
