package qian_conference

import (
	"dispatchAst/internal/consts"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

// 定义初始化方法
func New() *Controller {
	return &Controller{}
}

/**
 * 功能: 敬请期待
 *
 */
func (c *Controller) Get(req *ghttp.Request) {
	code, msg := 200, ""
	data := make(map[any]string, 0)

	room := req.GetQuery("room", "00").String()
	extension := req.GetQuery("extension", "00").String()

	if room != "00" {
		data[room] = consts.AA
		// TODO

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data[room],
		})
	}

	if extension != "00" {
		data[extension] = consts.BB
		// TODO

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data[extension],
		})
	}
}
