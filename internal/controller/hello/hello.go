package hello

import "github.com/gogf/gf/v2/net/ghttp"

type Hello struct{}

// 定义初始化方法
func New() *Hello {
	return &Hello{}
}

func (c *Hello) Laugh(req *ghttp.Request) {
	req.Response.Writeln("仟略研发...")
}
