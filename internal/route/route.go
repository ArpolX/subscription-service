package route

import (
	"subscription-service/internal/controller"

	_ "subscription-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi"
)

func Handlers(ctrl controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/subscription", func(r chi.Router) {
		r.Post("/create", ctrl.CreateSubscription)
		r.Put("/update", ctrl.UpdateSubscription)
		r.Delete("/delete", ctrl.DeleteSubscription)
		r.Get("/get", ctrl.GetIdSubscription)
		r.Get("/getList", ctrl.GetListSubscription)
	})

	return r
}
