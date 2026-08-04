package health

import (
	"github.com/Cheeginator/coldstart/service/health"
	"github.com/Cheeginator/coldstart/service/health/healthintf"
	"go.uber.org/fx"
)

type HandlerParams struct {
	fx.In

	HealthService *health.Service
}

type Handler struct {
	healthService healthintf.Service
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		healthService: params.HealthService,
	}
}
