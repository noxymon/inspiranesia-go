package echo

import (
	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	echo.Context
}

func newContext(ctx echo.Context) *EchoContext {
	return &EchoContext{ctx}
}

func (e EchoContext) GetNougatContextName() {
}