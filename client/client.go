package client

import (
	"github.com/Cheeginator/coldstart/client/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"client",
	fx.Provide(
		health.NewClient,
	),
)
