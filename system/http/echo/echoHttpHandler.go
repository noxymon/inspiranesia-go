package echo

import (
	"github.com/labstack/echo/v4"
	"inspiranesia/system/config"
	"inspiranesia/system/logging"
	"net/http"
	"strconv"
)

type EchoHttpHandler struct {
	port        int
	rootName    string
	httpHandler *echo.Echo
	logger      logging.NougatLoggingProvider
}

func New(server config.Server, logger logging.NougatLoggingProvider) *EchoHttpHandler {
	return &EchoHttpHandler{
		port:        server.Port,
		rootName:    server.RootName,
		httpHandler: echo.New(),
		logger:      logger,
	}
}

func (e EchoHttpHandler) Start() {
	e.httpHandler.GET(e.rootName, e.defaultRootEndpoint)
	e.logger.Fatal(e.httpHandler.Start(":" + strconv.Itoa(e.port)))
}

func (e EchoHttpHandler) defaultRootEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "This is default end point using echo :)")
}
