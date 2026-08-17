package models

import "time"

type User struct {
	ID             int64
	Email          string
	PasswordHash   string
	ActiveTenantID int64
	CreatedAt      time.Time
}
