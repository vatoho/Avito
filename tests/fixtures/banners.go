package fixtures

import (
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/tests/states"
)

type BannerBuilder struct {
	instance *model.Banner
}

func Banner() *BannerBuilder {
	return &BannerBuilder{instance: &model.Banner{}}
}

func (b *BannerBuilder) ID(v uint64) *BannerBuilder {
	b.instance.ID = v
	return b
}

func (b *BannerBuilder) TagIDs(v []uint64) *BannerBuilder {
	b.instance.TagIDs = v
	return b
}

func (b *BannerBuilder) FeatureID(v uint64) *BannerBuilder {
	b.instance.FeatureID = v
	return b
}

func (b *BannerBuilder) Content(v string) *BannerBuilder {
	b.instance.Content = v
	return b
}

func (b *BannerBuilder) CreatedAt(v time.Time) *BannerBuilder {
	b.instance.CreatedAt = v
	return b
}

func (b *BannerBuilder) UpdatedAt(v time.Time) *BannerBuilder {
	b.instance.UpdatedAt = v
	return b
}
func (b *BannerBuilder) IsActive(v bool) *BannerBuilder {
	b.instance.IsActive = v
	return b
}

func (b *BannerBuilder) Ptr() *model.Banner {
	return b.instance
}

func (b *BannerBuilder) Val() model.Banner {
	return *b.instance
}

func (b *BannerBuilder) Valid1() *BannerBuilder {
	return Banner().ID(states.ID1).TagIDs([]uint64{states.TagID1, states.TagID2}).FeatureID(states.FeatureID1).Content(states.Content1).
		CreatedAt(time.Time{}.Add(time.Hour)).UpdatedAt(time.Time{}.Add(time.Hour)).IsActive(states.IsActive1)
}

func (b *BannerBuilder) Valid2() *BannerBuilder {
	return Banner().ID(states.ID2).TagIDs([]uint64{states.TagID1, states.TagID4}).FeatureID(states.FeatureID2).Content(states.Content2).
		CreatedAt(time.Time{}.Add(time.Hour * 2)).UpdatedAt(time.Time{}.Add(time.Hour * 2)).IsActive(states.IsActive2)
}
