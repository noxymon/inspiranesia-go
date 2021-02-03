package http

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/labstack/echo/v4"
	"inspiranesia/system/config"
)

type NougatHttpHandlerAdapter struct {
	httpHandler NougatHttpHandler
	config      *config.DefaultConfig
}

type NougatHttpResponse struct {
	StatusCode int
	Body       interface{}
}

type handlerFunction func(param ...interface{}) *NougatHttpResponse

func ProvideHttpAdapter(handler NougatHttpHandler, defaultConfig *config.DefaultConfig) *NougatHttpHandlerAdapter {
	return &NougatHttpHandlerAdapter{
		httpHandler: handler,
		config:      defaultConfig,
	}
}

func (n NougatHttpHandlerAdapter) GetAndReturnStringF(endpointName string, h handlerFunction) {
	if n.config.Server.HttpHandler == "echo" {
		e := n.httpHandler.GetHttpHandler().(*echo.Echo)
		e.GET(endpointName, func(context echo.Context) error {
			param := context.QueryParam("param1")
			httpResponse := h(param)
			return context.String(httpResponse.StatusCode, httpResponse.Body.(string))
		})
	}
	e := n.httpHandler.GetHttpHandler().(*iris.Application)
	e.Get(endpointName, func(context context.Context) {
		param := context.URLParam("param1")
		response := h(param)
		context.StatusCode(response.StatusCode)
		context.WriteString(response.Body.(string))
	})
}

func (n NougatHttpHandlerAdapter) GetAndReturnJson(endpointName string, response NougatHttpResponse) {
	if n.config.Server.HttpHandler == "echo" {
		e := n.httpHandler.GetHttpHandler().(*echo.Echo)
		e.GET(endpointName, func(context echo.Context) error {
			return context.JSON(response.StatusCode, response.Body)
		})
	}
}
