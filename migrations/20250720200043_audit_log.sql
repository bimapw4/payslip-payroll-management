-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS audit_log (
    id              UUID            NOT NULL    PRIMARY KEY,
    user_id         VARCHAR         NOT NULL,
    request_id      VARCHAR         NOT NULL,
    ip_address      VARCHAR         NOT NULL,
    path_invoice    TEXT            NOT NULL,
    payload         JSONB           NOT NULL    DEFAULT '{}',
    response        JSONB           NOT NULL    DEFAULT '{}',
    created_at      TIMESTAMPTZ     NOT NULL    DEFAULT now(),
    updated_at      TIMESTAMPTZ     NOT NULL    DEFAULT now(),
    created_by      VARCHAR         NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS audit_log;
-- +goose StatementEnd
