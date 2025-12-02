package repository

import (
	"context"
	"subscription-service/internal/entity"
	db "subscription-service/internal/infrastructure/database/postgres"

	"go.uber.org/zap"
)

type Repository interface {
	CreateSubscription(ctx context.Context, s entity.SubscriptionRequest) error
	UpdateSubscription(ctx context.Context, s entity.SubscriptionRequest) error
	DeleteSubscription(ctx context.Context, id string) error
	GetSubscription(ctx context.Context, id string) (entity.SubscriptionResponse, error)
	GetListSubscription(ctx context.Context) ([]entity.SubscriptionResponse, error)
}

type RepositoryImpl struct {
	Log      *zap.Logger
	Postgres *db.Postgres
}

func NewRepositoryImpl(postgres *db.Postgres, log *zap.Logger) Repository {
	return &RepositoryImpl{
		Log:      log,
		Postgres: postgres,
	}
}
