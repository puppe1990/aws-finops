package models

import "time"

type Tenant struct {
	ID        int64
	Name      string
	Slug      string
	CreatedAt time.Time
}

type Membership struct {
	TenantID int64
	UserID   int64
	Role     string
}
