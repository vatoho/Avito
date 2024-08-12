package service

import (
	"context"
)

func (s *BannerServiceApp) DeleteBanner(ctx context.Context, id uint64) error {
	return s.storage.DeleteBanner(ctx, id)
}
