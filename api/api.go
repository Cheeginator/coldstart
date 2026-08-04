package api

import (
	"github.com/Cheeginator/coldstart/api/handler"
	"github.com/Cheeginator/coldstart/api/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	fx.Invoke(
		server.NewServer,
	),
	handler.Module,
)
