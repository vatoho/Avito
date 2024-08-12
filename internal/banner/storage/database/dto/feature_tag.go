package dto

type FeatureTag struct {
	FeatureID uint64 `db:"feature_id"`
	TagID     uint64 `db:"tag_id"`
}
