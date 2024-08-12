package model

import "time"

type BannerVersion struct {
	VersionID uint64
	BannerID  uint64
	Content   string
	UpdatedAt time.Time
	IsActive  bool
}
