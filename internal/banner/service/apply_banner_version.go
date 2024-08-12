package service

import "context"

func (s *BannerServiceApp) ApplyBannerVersion(ctx context.Context, versionID uint64) error {
	return s.storage.ApplyBannerVersion(ctx, versionID)
}
