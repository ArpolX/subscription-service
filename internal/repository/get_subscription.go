package repository

import (
	"context"
	"fmt"
	"subscription-service/internal/entity"
	"time"
)

const qGetSub = `select service_name, price, user_id, start_date, end_date from subscription where subscription_id = $1`

func (r *RepositoryImpl) GetSubscription(ctx context.Context, id string) (entity.SubscriptionResponse, error) {
	var s entity.SubscriptionResponse
	var start, end *time.Time

	err := r.Postgres.DB.QueryRow(ctx, qGetSub, id).Scan(
		&s.ServiceName,
		&s.Price,
		&s.UserId,
		&start,
		&end,
	)
	if err != nil {
		return entity.SubscriptionResponse{}, fmt.Errorf("ошибка в query, метод GetSubscription: %w", err)
	}

	s.StartDate = start
	s.EndDate = end

	return s, nil
}
