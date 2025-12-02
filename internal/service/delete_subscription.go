package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

func (s *ServiceImpl) DeleteSubscription(ctx context.Context, id string) error {
	if err := s.Repo.DeleteSubscription(ctx, id); err != nil {
		s.Log.Error("ошибка при удалении подписки, метод DeleteSubscription", zap.Error(err))
		return fmt.Errorf("ошибка при удалении подписки, метод DeleteSubscription: %w", err)
	}

	return nil
}
