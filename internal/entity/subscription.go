package entity

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionRequest struct {
	SubscriptionId string     `json:"subscription_id"`
	ServiceName    string     `json:"service_name"`
	Price          int        `json:"price"`
	UserId         *uuid.UUID `json:"user_id"`
	StartDate      *string    `json:"start_date"`
	EndDate        *string    `json:"end_date"`
}

type SubscriptionResponse struct {
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserId      string     `json:"user_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}
