package repository

import (
	"context"
	"fmt"
	"subscription-service/internal/entity"
	"subscription-service/pkg/errors"
)

const qCreateSub = `Insert into subscription
							(subscription_id, service_name, price, user_id, start_date, end_date)
							values ($1, $2, $3, $4, $5, $6)`

func (r *RepositoryImpl) CreateSubscription(ctx context.Context, s entity.SubscriptionRequest) error {
	_, err := r.Postgres.DB.Exec(ctx, qCreateSub, s.SubscriptionId, s.ServiceName, s.Price, s.UserId, s.StartDate, s.EndDate)
	if err != nil {
		if IsUniqueViolation(err) {
			return fmt.Errorf("%w", errors.ID_EXISTS)
		}
		return fmt.Errorf("ошибка в Exec, метод CreateSubscription: %w", err)
	}

	return nil
}
