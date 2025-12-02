package controller

import (
	"errors"
	"fmt"
	"net/http"
	"subscription-service/internal/entity"
	Error "subscription-service/pkg/errors"

	"go.uber.org/zap"
)

// @Summary Регистрация новой подписки
// @Description Создание новой подписки, subscription_id - primary key. user_id и start_date можно не указывать, на них стоит default. end_date также можно не указывать, в бд допускается null
// @Tags subscription
// @Accept json
// @Produce plain
// @Param subscription body entity.SubscriptionRequest false "Заполните поля согласно описанию"
// @Success 200 {string} Info "Успешное создание"
// @Failure 400 {object} entity.ErrorResponse "Ошибка"
// @Failure 500 {object} entity.ErrorResponse "Ошибка"
// @Router /subscription/create [post]
func (c *ControllerImpl) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var s entity.SubscriptionRequest

	if err := decoder.Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c.Log.Error("Ошибка обработки пути /subscription/create, метод CreateSubscription", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	if err := c.Srv.CreateSubscription(ctx, s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if errors.Is(err, Error.ID_EXISTS) {
			CreateError("ID_EXISTS", Error.ID_EXISTS.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /subscription/create, метод CreateSubscription", zap.Error(err))
		CreateError("500", err.Error(), w)
		return
	}

	w.Write([]byte("Подписка создана"))
}
