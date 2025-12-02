package controller

import (
	"errors"
	"net/http"
	_ "subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

// @Summary Получение подписки
// @Description Получение подписки по subscription_id
// @Tags subscription
// @Produce json
// @Param subscription_id query string true "id подписки"
// @Success 200 {object} entity.SubscriptionResponse "Успешное получение подписки"
// @Failure 400 {object} entity.ErrorResponse "Ошибка"
// @Router /subscription/get [get]
func (c *ControllerImpl) GetIdSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("subscription_id")

	sResp, err := c.Srv.GetIdSubscription(ctx, id)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", Error.NOT_FOUND.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /subscription/get, метод GetIdSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&sResp); err != nil {
		c.Log.Error("Ошибка обработки пути /subscription/get, метод GetIdSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
