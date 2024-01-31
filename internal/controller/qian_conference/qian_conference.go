package qian_conference

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

// bridge类型的会议
type ConfBridge struct {
	BridgeName  string `json:"name"`
	UsersNumber int    `json:"users_number"`
	Marked      int    `json:"marked"`
	Locked      string `json:"locked"` // bool
	Muted       string `json:"muted"`  // bool
}

// 定义初始化方法
func New() *Controller {
	return &Controller{}
}

/**
 * 功能: 敬请期待
 *
 * e.g ip:port/qianlue/api/meet?rooms=1
 */
func (c *Controller) Get(req *ghttp.Request) {
	code, msg := 200, ""

	rooms := req.GetQuery("rooms", "00").String()
	room := req.GetQuery("room", "00").String()
	extension := req.GetQuery("extension", "00").String()

	if rooms != "00" {
		var data []ConfBridge

		if count := GetConfCount(); count == 0 {
			msg, data = "当前没有会议", nil
		} else {
			msg, data = fmt.Sprintf("正在进行 %d 场会议", count), GetRooms()
		}

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
	}

	// 获取单个房间里参与者信息
	if room != "00" {
		var data []any
		// TODO

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
	}

	if extension != "00" {
		var data []any
		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
	}
}
