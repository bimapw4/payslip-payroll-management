-- +goose Up
-- +goose StatementBegin
ALTER TABLE reimbursement DROP COLUMN IF EXISTS attachment;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reimbursement ADD COLUMN IF NOT EXISTS attachment TEXT;
-- +goose StatementEnd
