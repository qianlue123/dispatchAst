package qian_force

import (
	"context"
	"dispatchAst/api/qian_force"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

// 定义初始化方法
func New() *Controller {
	return &Controller{}
}

func (c *Controller) Params(ctx context.Context, req *qian_force.ParamsReq) (res *qian_force.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.Writeln("params")
	name := r.GetQuery("name")

	r.Response.Writeln(name)
	return
}

func (c *Controller) Db(req *ghttp.Request) {
	req.Response.Writeln(g.Model("users"))
}
