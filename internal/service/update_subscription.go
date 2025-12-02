package service

import (
	"context"
	"fmt"
	"subscription-service/internal/entity"
	defaultformat "subscription-service/pkg/default_format"

	"go.uber.org/zap"
)

func (s *ServiceImpl) UpdateSubscription(ctx context.Context, sReq entity.SubscriptionRequest) error {
	sReq = defaultformat.DefaultFormat(sReq)

	if err := s.Repo.UpdateSubscription(ctx, sReq); err != nil {
		s.Log.Error("ошибка при обновлении подписки, метод UpdateSubscription", zap.Error(err))
		return fmt.Errorf("ошибка при обновлении подписки, метод UpdateSubscription: %w", err)
	}

	return nil
}
