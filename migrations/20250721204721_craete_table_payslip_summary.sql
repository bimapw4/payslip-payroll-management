-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payslip_summary (
    id                  UUID        NOT NULL    PRIMARY KEY,
    payroll_id          UUID        NOT NULL    REFERENCES payroll(id) ON DELETE CASCADE,
    user_id             UUID        NOT NULL    REFERENCES users(id) ON DELETE CASCADE,
    base_salary         INT         NOT NULL,
    prorated_salary     INT         NOT NULL,
    overtime_pay        INT         NOT NULL,
    reimbursement_total INT         NOT NULL,
    take_home_pay       INT         NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL    DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL    DEFAULT now(),
    created_by          VARCHAR     NOT NULL,
    updated_by          VARCHAR     NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payslip_summary;
-- +goose StatementEnd
