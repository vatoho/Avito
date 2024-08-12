package model

import "time"

type Banner struct {
	ID        uint64
	TagIDs    []uint64
	FeatureID uint64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
}
