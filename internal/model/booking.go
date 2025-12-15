package model

import (
	"database/sql"
	"time"
)

type BookingData struct {
	Nama         sql.NullString
	Nim          sql.NullString
	Schedule     time.Time
	UniversityID int64
}
