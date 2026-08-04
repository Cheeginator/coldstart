package handler

import (
	"github.com/Cheeginator/coldstart/api/handler/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handler",
	fx.Provide(
		health.NewHandler,
	),
)
