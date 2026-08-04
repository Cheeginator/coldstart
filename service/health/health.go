package health

import (
	"github.com/Cheeginator/coldstart/client/health"
	"github.com/Cheeginator/coldstart/client/health/healthintf"
	healthe "github.com/Cheeginator/coldstart/entity/health"
	healtheintf "github.com/Cheeginator/coldstart/entity/health/healthintf"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	HealthClient *health.Client
	HealthEntity *healthe.Entity
}

type Service struct {
	healthClient healthintf.Client
	healthEntity healtheintf.Entity
}

func NewService(params ServiceParams) *Service {
	return &Service{
		healthClient: params.HealthClient,
		healthEntity: params.HealthEntity,
	}
}
