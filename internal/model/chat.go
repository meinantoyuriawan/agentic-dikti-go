package model

import (
	"database/sql"
	"time"
)

type ChatHistory struct {
	Role      sql.NullString
	ChatInput sql.NullString
}

type ChatLogs struct {
	SessionID    sql.NullString
	ChatID       sql.NullString
	ChatInput    sql.NullString
	Timestamp    time.Time
	Role         sql.NullString
	Emergency    bool
	UniversityID int64
}
