package controller

import (
	"errors"
	"net/http"
	_ "subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

// @Summary Удаление подписки
// @Description Удаление подписки по subscription_id
// @Tags subscription
// @Produce plain
// @Param subscription_id query string true "id подписки"
// @Success 200 {string} Info "Успешное удаление"
// @Failure 400 {object} entity.ErrorResponse "Ошибка"
// @Router /subscription/delete [delete]
func (c *ControllerImpl) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("subscription_id")

	if err := c.Srv.DeleteSubscription(ctx, id); err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", Error.NOT_FOUND.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /subscription/delete, метод DeleteSubscription", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	w.Write([]byte("Подписка удалена"))
}
