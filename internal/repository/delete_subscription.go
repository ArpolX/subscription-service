package repository

import (
	"context"
	"fmt"
	Error "subscription-service/pkg/errors"
)

const qDeleteSub = `Delete from subscription where subscription_id = $1`

func (r *RepositoryImpl) DeleteSubscription(ctx context.Context, id string) error {
	tag, err := r.Postgres.DB.Exec(ctx, qDeleteSub, id)
	if err != nil {
		return fmt.Errorf("ошибка в Delete, метод DeleteSubscription: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%w", Error.NOT_FOUND)
	}

	return nil
}
