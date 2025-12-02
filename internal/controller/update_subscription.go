package controller

import (
	"errors"
	"fmt"
	"net/http"
	"subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

// @Summary Обновление подписки
// @Description Обновление подписки, subscription_id - primary key. service_name, user_id не учитывается. start_date можно не указывать, стоит default. end_date также можно не указывать, в бд будет null
// @Tags subscription
// @Accept json
// @Produce plain
// @Param subscription body entity.SubscriptionRequest false "service_name, user_id не учитывается. start_date можно не указывать"
// @Success 200 {string} Info "Успешное обновление"
// @Failure 400 {object} entity.ErrorResponse "Ошибка"
// @Router /subscription/update [put]
func (c *ControllerImpl) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var s entity.SubscriptionRequest

	if err := decoder.Decode(&s); err != nil {
		c.Log.Error("Ошибка обработки пути /subscription/update, метод UpdateSubscription", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	if err := c.Srv.UpdateSubscription(ctx, s); err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", Error.NOT_FOUND.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /subscription/update, метод UpdateSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	w.Write([]byte("Подписка обновлена"))
}
