package dto

import "github.com/ilyushkaaa/banner-service/internal/banner/model"

func ConvertToBanner(bannerDB BannerDB) model.Banner {
	return model.Banner{
		ID:        bannerDB.ID,
		FeatureID: bannerDB.FeatureID,
		Content:   bannerDB.Content,
		CreatedAt: bannerDB.CreatedAt,
		UpdatedAt: bannerDB.UpdatedAt,
		IsActive:  bannerDB.IsActive,
	}
}

func ConvertToBannerSlice(bannersDB []BannerDB) []model.Banner {
	banners := make([]model.Banner, 0, len(bannersDB))
	for _, b := range bannersDB {
		banners = append(banners, ConvertToBanner(b))
	}
	return banners
}

func ConvertToBannerVersion(bannerDB BannerVersionDB) model.BannerVersion {
	return model.BannerVersion{
		VersionID: bannerDB.VersionID,
		BannerID:  bannerDB.BannerID,
		Content:   bannerDB.Content,
		UpdatedAt: bannerDB.UpdatedAt,
		IsActive:  bannerDB.IsActive,
	}
}

func ConvertToBannerVersionSlice(bannersDB []BannerVersionDB) []model.BannerVersion {
	banners := make([]model.BannerVersion, 0, len(bannersDB))
	for _, b := range bannersDB {
		banners = append(banners, ConvertToBannerVersion(b))
	}
	return banners
}
