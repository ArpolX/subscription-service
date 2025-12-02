package controller

import (
	"errors"
	"net/http"
	_ "subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

// @Summary Получение всех подписок
// @Description Получение всех подписок
// @Tags subscription
// @Produce json
// @Success 200 {object} []entity.SubscriptionResponse "Успешное получение подписки"
// @Failure 400 {object} entity.ErrorResponse "Ошибка"
// @Router /subscription/getList [get]
func (c *ControllerImpl) GetListSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sRespList, err := c.Srv.GetListSubscription(ctx)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", Error.NOT_FOUND.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /subscription/getList, метод GetListSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&sRespList); err != nil {
		c.Log.Error("Ошибка обработки пути /subscription/getList, метод GetListSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
