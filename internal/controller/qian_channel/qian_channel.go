package qian_channel

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

// 定义初始化方法
func New() *Controller {
	return &Controller{}
}

/*
e.g ip:port/qianlue/api/channel?bridges=1
params usage:
  - statistics  针对统计方面的请求, 专门返回单个数值
  - bridges     专门返回对象
*/
func (c *Controller) Get(req *ghttp.Request) {
	code, msg := 200, ""

	statistics, bridges, bridge := req.GetQuery("statistics", "00").String(),
		req.GetQuery("bridges", "00").String(),
		req.GetQuery("bridge", "00").String()

	if statistics != "00" {
		data := getValue4Statistics(statistics)

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
		return
	}

	if bridges != "00" {
		bridgeNameArr := GetBridgeName()

		data := bridgeNameArr

		if bridges == "1" {
			req.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    code,
				Message: msg,
				Data:    data,
			})
			return
		}
	}

	// 提供具体桥接名
	if bridge != "00" {
		bOne := NewBridge(bridge)

		fmt.Println("one:", bOne)

		exts := bOne.ShowExtension()
		fmt.Println("exts in birdge: ", exts)

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    bOne,
		})
		return
	}

	bridgeNameArr := GetBridgeName()
	Bridges := make([]Bridge, 0)

	for i := 0; i < len(bridgeNameArr); i++ {
		bOne := Bridge{}
		bOne.SetName(bridgeNameArr[i])
		bOne.SetType(bridgeNameArr[i])

		channelNameArr := bOne.ShowChannelName()

		for j := 0; j < len(channelNameArr); j++ {
			cOne := NewChannel(channelNameArr[j])

			bOne.channels = append(bOne.channels, *cOne)
		}

		Bridges = append(Bridges, bOne)
	}

	fmt.Println("Bridges: ", Bridges)

}

func getValue4Statistics(what string) int {
	data := 0

	switch what {
	case "bridges":
		bridgeNameArr := GetBridgeName()
		data = len(bridgeNameArr)

	default:
	}

	return data
}
