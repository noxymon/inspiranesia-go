package http

import (
	"inspiranesia/system/config"
	"inspiranesia/system/http/echo"
	"inspiranesia/system/http/iris"
	"inspiranesia/system/logging"
)

type NougatHttpHandler interface {
	Start()
	GetHttpHandler() interface{}
}

func ProvideHttpHandler(config *config.DefaultConfig, logger logging.NougatLoggingProvider) NougatHttpHandler {
	serverConfig := config.Server
	if serverConfig.HttpHandler == "echo" {
		return echo.New(serverConfig, logger)
	}
	return iris.New(serverConfig, logger)
}
