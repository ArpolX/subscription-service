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

func (s *ServiceImpl) GetListSubscription(ctx context.Context) ([]entity.SubscriptionResponse, error) {
	sRespList, err := s.Repo.GetListSubscription(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("ошибка при получении подписок, метод GetListSubscription", zap.Error(err))
		return nil, fmt.Errorf("ошибка при получении подписок, метод GetListSubscription: %w", err)
	}

	return sRespList, nil
}
