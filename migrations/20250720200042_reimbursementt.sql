-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reimbursement (
    id          UUID        NOT NULL PRIMARY KEY,
    user_id     UUID        ,
    amount      INT         NOT NULL,
    description TEXT        NOT NULL,
    attachment  TEXT        NOT NULL,
    payroll_id  VARCHAR     ,
    created_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by  VARCHAR     NOT NULL,
    updated_by  VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reimbursement;
-- +goose StatementEnd
