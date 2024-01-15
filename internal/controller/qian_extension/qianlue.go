package qian_extension

import (
	"context"
	"dispatchAst/api/qian_info"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 定义枚举映射话机的状态
const (
	Idle     = iota
	NotInUse = 0
	InUse    = iota
	Ring
	Ringing
	Unavailable
	_
	ALL
)

type Controller struct{}

// 定义初始化方法
func New() *Controller {
	return &Controller{}
}

// 话机信息
type extension struct {
	CID     string
	Desc    string // 用户名
	ExtName string // 电话号
	ExtAge  int
	ExtPwd  string
}

/**
 * 功能: 专门用来查所有相同状态的话机集合
 *
 * e.g ip:port/qianlue/api/params?extensions=inuse
 */
func (c *Controller) Params(ctx context.Context, req *qian_info.ParamsReq) (res *qian_info.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)

	code, msg, extensions := 200, "", make([]extension, 0)
	msg = "list qianlue data, all extension whose who are "

	// url 中没有提供值, 当作查所有可用状态的话机
	state := r.GetQuery("extensions", "idle")

	// 捕捉可能出现的输入情况, 当作同一类处理
	switch state.String() {
	case "InUse", "Inuse", "inuse", "in use":
		extensions = GetAllExtension(InUse)
		msg += "in use"

	case "Ringing", "ringing":
		extensions = GetAllExtension(Ringing)
		msg += "ringing"

	case "Ring", "ring":
		extensions = GetAllExtension(Ring)
		msg += "ring"

	case "Unavailable", "unavailable":
		// 所有注册过但是没人用了的电话, 一般用不到
		extensions = GetAllExtension(Unavailable)
		msg += "unavailable"

	case "Idle", "idle":
		// 包含默认值
		fallthrough
	case "NotInUse", "Notinuse", "notinuse", "not in use":
		// 所有可用电话
		extensions = GetAllExtension(NotInUse)
		msg += "not in use"

	default:
		code = 0
		msg = "参数传入错误"
	}

	if err := r.Parse(&extensions); err != nil {
		fmt.Println(err)
	}
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    extensions,
	})

	return
}

// 响应输出
func (c *Controller) Respons(ctx context.Context, req *qian_info.ParamsReq) (res *qian_info.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.Writeln("ext")
	r.Response.WritelnExit("ext") // 页面显示完后面不显示
	return
}

func (c *Controller) Add(req *ghttp.Request) {
	req.Response.Writeln("添加话机")
}

func (c *Controller) Update(req *ghttp.Request) {
	req.Response.Writeln("更新话机")
}

func (c *Controller) Delete(req *ghttp.Request) {
	req.Response.Writeln("删除话机")
}

func (c *Controller) Get(req *ghttp.Request) {
	GetStateExt(ALL)

	req.Response.Writeln("查询话机接口")

	info := ""
	if info = GetStateExt(InUse); info != "0" {
		req.Response.Writeln("正在通话中\t", info)
	} else {
		req.Response.Writeln("没人用电话")
	}

	if info = GetStateExt(Ring); info != "0" {
		req.Response.Writeln("正在呼叫的电话\n", info)
	} else {
		req.Response.Writeln("没有电话呼叫")
	}

	if info = GetStateExt(Ringing); info != "0" {
		req.Response.Writeln("正在振铃的电话\n", info)
	} else {
		req.Response.Writeln("没有电话振铃")
	}
}

func (c *Controller) Put(req *ghttp.Request) {
	req.Response.Writeln("修改话机")
}

func (c *Controller) DisplayInfo(req *ghttp.Request) {
	req.Response.Writeln("一切就绪...")
}
