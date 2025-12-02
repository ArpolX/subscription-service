package service

import (
	"context"
	"subscription-service/internal/entity"
	"subscription-service/internal/repository"

	"go.uber.org/zap"
)

type Service interface {
	CreateSubscription(ctx context.Context, s entity.SubscriptionRequest) error
	UpdateSubscription(ctx context.Context, sReq entity.SubscriptionRequest) error
	DeleteSubscription(ctx context.Context, id string) error
	GetIdSubscription(ctx context.Context, id string) (entity.SubscriptionResponse, error)
	GetListSubscription(ctx context.Context) ([]entity.SubscriptionResponse, error)
}

type ServiceImpl struct {
	Log  *zap.Logger
	Repo repository.Repository
}

func NewServiceImpl(log *zap.Logger, repo repository.Repository) Service {
	return &ServiceImpl{
		Log:  log,
		Repo: repo,
	}
}
