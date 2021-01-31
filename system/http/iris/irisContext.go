package iris

import (
	"github.com/kataras/iris/v12/context"
)

type IrisContext struct {
	context.Context
}

func newContext(ctx context.Context) *IrisContext {
	return &IrisContext{ctx}
}

func (i IrisContext) GetNougatContextName() {}