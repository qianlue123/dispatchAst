package qian_extension

import (
	"context"
	"dispatchAst/api/qian_info"

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
	ExtName string
	ExtAge  int
	ExtPwd  string
}

// ip:port/xxx/params?name=xxx
func (c *Controller) Params(ctx context.Context, req *qian_info.ParamsReq) (res *qian_info.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)

	// uri 中没有 name 或没用提供值, 用 qianlue 代替
	//name := r.GetQuery("name", "qianlue")

	//r.Response.Writeln(name.String() + "good!")

	// 对应 <input> 里的 name 属性
	var ext extension
	r.Parse(&ext)
	r.Response.Writeln(ext)

	// 所有在用电话
	extensionsInUse := GetAllExtension(InUse)
	r.Response.WriteJson(extensionsInUse)

	// 所有可用电话
	extensionsNotInUse := GetAllExtension(NotInUse)
	r.Response.WriteJson(extensionsNotInUse)

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
