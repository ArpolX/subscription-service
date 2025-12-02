package repository

import (
	"context"
	"fmt"
	"subscription-service/internal/entity"
	"subscription-service/pkg/errors"
)

const qUpdateSub = `update subscription
							set price = $1, start_date = $2, end_date = $3
							where subscription_id = $4`

func (r *RepositoryImpl) UpdateSubscription(ctx context.Context, s entity.SubscriptionRequest) error {
	tag, err := r.Postgres.DB.Exec(ctx, qUpdateSub, s.Price, s.StartDate, s.EndDate, s.SubscriptionId)
	if err != nil {
		return fmt.Errorf("ошибка в exec, метод UpdateSubscription: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%w", errors.NOT_FOUND)
	}

	return nil
}
