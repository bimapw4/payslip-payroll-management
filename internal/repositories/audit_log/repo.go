package auditlog

import "github.com/jmoiron/sqlx"

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) AuditLog {
	return &repo{
		db: db,
	}
}
