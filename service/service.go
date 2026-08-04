package service

import (
	"github.com/Cheeginator/coldstart/service/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Provide(
		health.NewService,
	),
)
