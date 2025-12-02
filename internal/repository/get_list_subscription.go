package repository

import (
	"context"
	"fmt"
	"subscription-service/internal/entity"
	"time"
)

const qGetListSub = `select service_name, price, user_id, start_date, end_date from subscription`

func (r *RepositoryImpl) GetListSubscription(ctx context.Context) ([]entity.SubscriptionResponse, error) {
	rows, err := r.Postgres.DB.Query(ctx, qGetListSub)
	if err != nil {
		return nil, fmt.Errorf("ошибка в Query, метод GetListSubscription: %w", err)
	}
	defer rows.Close()

	var sList []entity.SubscriptionResponse

	for rows.Next() {
		var s entity.SubscriptionResponse
		var start, end *time.Time

		err := rows.Scan(
			&s.ServiceName,
			&s.Price,
			&s.UserId,
			&start,
			&end,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка в rows, метод GetListSubscription: %w", err)
		}

		s.StartDate = start
		s.EndDate = end

		sList = append(sList, s)
	}

	return sList, nil
}
