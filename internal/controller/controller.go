package controller

import (
	"net/http"
	"subscription-service/internal/config"
	"subscription-service/internal/service"

	"go.uber.org/zap"

	jsonIterator "github.com/json-iterator/go"
)

var Json = jsonIterator.ConfigCompatibleWithStandardLibrary

type Controller interface {
	CreateSubscription(w http.ResponseWriter, r *http.Request)
	UpdateSubscription(w http.ResponseWriter, r *http.Request)
	DeleteSubscription(w http.ResponseWriter, r *http.Request)
	GetIdSubscription(w http.ResponseWriter, r *http.Request)
	GetListSubscription(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	Log *zap.Logger
	Srv service.Service
}

func NewControllerImpl(cfg config.Config, log *zap.Logger, srv service.Service) ControllerImpl {
	return ControllerImpl{
		Log: log,
		Srv: srv,
	}
}
