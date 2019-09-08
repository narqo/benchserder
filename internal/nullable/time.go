package nullable

import (
	"database/sql"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not null
}

func NewTime(value time.Time) Time {
	if value.IsZero() {
		return Time{}
	}
	return Time{
		Time:  value,
		Valid: true,
	}
}

func (v Time) IsDefined() bool {
	return v.Valid
}

func (v Time) Format(layout string) sql.NullString {
	if !v.Valid {
		return sql.NullString{}
	}
	return sql.NullString{
		Valid:  true,
		String: v.Time.Format(layout),
	}
}
