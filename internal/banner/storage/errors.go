package storage

import "errors"

var (
	ErrBannerNotFound      = errors.New("no banners with such id")
	ErrDuplicateFeatureTag = errors.New("such pair of feature and tag already exists")
	ErrNoBannerVersions    = errors.New("no banner versions with such id")
)
