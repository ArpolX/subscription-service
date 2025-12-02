package defaultformat

import (
	"subscription-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

func DefaultFormat(subscription entity.SubscriptionRequest) entity.SubscriptionRequest {
	if subscription.UserId == nil {
		id := uuid.New()
		subscription.UserId = &id
	}

	if subscription.StartDate == nil {
		now := time.Now().Format(time.RFC3339)
		subscription.StartDate = &now
	}

	return subscription
}
