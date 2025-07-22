package presentations

import (
	"encoding/json"
	"time"
)

type AuditLog struct {
	ID        string          `db:"id" json:"id"`
	UserID    string          `db:"user_id" json:"user_id"`
	RequestID string          `db:"request_id" json:"request_id"`
	IPAddress string          `db:"ip_address" json:"ip_address"`
	Path      string          `db:"path_invoice" json:"path"`
	Payload   json.RawMessage `db:"payload" json:"payload"`
	Response  json.RawMessage `db:"response" json:"response"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
	CreatedBy string          `db:"created_by" json:"created_by"`
}
