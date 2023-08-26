package common

import "time"

type Audit struct {
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func DefaultAudit() Audit {
	return Audit{
		CreatedAt: time.Now().In(time.UTC),
		CreatedBy: "system",
		UpdatedAt: time.Now().In(time.UTC),
		UpdatedBy: "system",
	}
}
