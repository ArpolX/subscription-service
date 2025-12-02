package service

import (
	"context"
	"errors"
	"fmt"
	"subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (s *ServiceImpl) GetIdSubscription(ctx context.Context, id string) (entity.SubscriptionResponse, error) {
	sResp, err := s.Repo.GetSubscription(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.SubscriptionResponse{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("ошибка при получении подписки, метод GetIdSubscription", zap.Error(err))
		return entity.SubscriptionResponse{}, fmt.Errorf("ошибка при получении подписки, метод GetIdSubscription: %w", err)
	}

	return sResp, nil
}
