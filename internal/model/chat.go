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
	SessionID    string
	ChatInput    string
	Timestamp    time.Time
	Role         string
	Emergency    bool
	UniversityID int
}
