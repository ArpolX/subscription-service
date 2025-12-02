package service

import (
	"context"
	"errors"
	"fmt"
	"subscription-service/internal/entity"
	defaultformat "subscription-service/pkg/default_format"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

func (s *ServiceImpl) CreateSubscription(ctx context.Context, subscription entity.SubscriptionRequest) error {
	subscription = defaultformat.DefaultFormat(subscription)

	if err := s.Repo.CreateSubscription(ctx, subscription); err != nil {
		if errors.Is(err, Error.ID_EXISTS) {
			return err
		}
		s.Log.Error("ошибка при создании подписки, метод CreateSubscription", zap.Error(err))
		return fmt.Errorf("ошибка при создании подписки, метод CreateSubscription: %w", err)
	}

	return nil
}
