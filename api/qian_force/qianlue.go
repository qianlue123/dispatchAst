package qian_force

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	g.Meta `method: "all"` // 什么样的请求都可以
}

type ParamsRes struct{}