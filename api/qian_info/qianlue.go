package qian_info

import "github.com/gogf/gf/v2/frame/g"

// 接受从网页传过来的参数
type ParamsReq struct {
	g.Meta `method: "all"` // 什么样的请求都可以
}

type ParamsRes struct{}
