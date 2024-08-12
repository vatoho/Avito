package dto

type BannerUpdate struct {
	TagIDs    []uint64 `json:"tag_ids"`
	FeatureID uint64   `json:"feature_id"`
	Content   string   `json:"content"`
	IsActive  bool     `json:"is_active"`
}
