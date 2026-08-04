package main

import (
	"github.com/Cheeginator/coldstart/api"
	"github.com/Cheeginator/coldstart/client"
	"github.com/Cheeginator/coldstart/config"
	"github.com/Cheeginator/coldstart/entity"
	"github.com/Cheeginator/coldstart/service"
	_ "github.com/Cheeginator/slogctx/sloginit"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		api.Module,
		client.Module,
		config.Module,
		entity.Module,
		service.Module,
	)
	app.Run()
}
