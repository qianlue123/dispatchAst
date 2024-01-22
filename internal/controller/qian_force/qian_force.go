package qian_force

import (
	"bytes"
	"context"
	"dispatchAst/api/qian_force"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

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

	tb_users := g.Model("users")
	m, _ := tb_users.One()
	req.Response.Writeln(m["extension"])

	// WriteJson 能直接将数据按 json 展示出来，但是和表中顺序不一样，会自动按首字母排序
	req.Response.WriteJson(m)

	// devices 表中记录了所有话机的信息
	tbDevices := g.Model("devices")
	extensions, _ := tbDevices.Fields("id", "user").All()
	req.Response.WriteJson(extensions)

	req.Response.Writeln(tb_users.Count())
}

/**
 * 功能: 获取使用中话机对应的 channel
 *
 * e.g ip:port/qianlue/force?extension=8008
 */
func (c *Controller) Get(req *ghttp.Request) {
	code, msg := 200, ""

	extension := req.GetQuery("extension", "0")
	statistics := req.GetQuery("statistics", "0")
	thing := req.GetQuery("dothing2", "0")

	// 有 extension 参数就做这件事
	if extension.String() != "0" {
		data := make([]Channel, 0)

		if activechannels, exist := getChannelCount(); exist {
			if activechannels != getCallCount()*2 {
				msg = "OK, but some channels comes from force plug!"
			}

			data = GetChannelInfo(extension.String())

			if err := req.Parse(&data); err != nil {
				code, msg = 1, "error: "+err.Error()
			}
		} else {
			code, msg = 1, "there is no Extension in use!"
		}

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
	}

	// 统计类查找, 返回的都是数字
	if statistics.String() != "0" {
		data := 0
		switch statistics.String() {
		case "calls":
			data = getCallCount()

		default:
		}

		req.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
	}

	if thing.String() != "0" {
		code, msg = 200, ""
		//TODO
	}
}

// 获取channel id
func (c *Controller) Post(req *ghttp.Request) {
	data := make(map[string]string, 1)
	data["channelid"] = ""
	json.NewDecoder(req.Body).Decode(&data)

	code, msg := 200, ""
	if len(data["channelid"]) == 0 {
		code, msg = 0, "客户端post请求出错, 提供的数据不符合要求"
	} else {
		fmt.Printf("%v", data)
		// channelID := req.GetQuery("channelid", "PJSIP/2024-00000066")
		channelID := data["channelid"]
		fmt.Println("channelid from client: ", channelID)

		cmd := fmt.Sprintf(asteriskrx[51], channelID)
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Printf("Failed to execute command: %s", cmd)
		}
		fmt.Println(string(out))
		code, msg = 200, "kill channel "+channelID+" ok"
	}
	req.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    "",
	})
}

// 更新多个呼叫的分机
func (c *Controller) Put(req *ghttp.Request) {
	code, msg, extNameArr := 200, "ok", make([]string, 0)
	defer req.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    extNameArr,
	})

	var info map[string]any
	json.NewDecoder(req.Body).Decode(&info)
	for _, params := range info {
		// fmt.Printf("%T, %v\n", v, v)
		if con, ok := params.([]any); ok {
			for _, element := range con {
				extNameArr = append(extNameArr, element.(string))
			}
		}
	}

	fmt.Println(extNameArr[0], extNameArr[1])
	// 此处不检测提供的分机是否存在, 由前端控制
	// 离线的分机不显示, 在通话中的分机还可以再呼叫

	cmd := fmt.Sprintf(asteriskrx[61], extNameArr[0], extNameArr[1])
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		code, msg = 1, "Failed to execute call them"
		return
	}

}

func GetChannelInfo(extName string) []Channel {
	var channels []Channel

	// awk '{OFS=":"} {print $1, $3}'
	cmd := mapBash[21] + "| awk '{OFS=\":\"} {print $1,$3}' "
	out, _ := exec.Command("bash", "-c", cmd).Output()

	out = out[:len(out)-1]
	outs := bytes.Split(out, []byte("\n"))
	for i, content := range outs {
		// 在格式 PJSIP/1234:Up 里提取 PJSIP/1234 当通道号, 提取 Up 当通道状态
		v := strings.Split(string(content), ":")
		if index := strings.Index(v[0], extName); index != -1 {
			fmt.Printf("%d %v \n", i, string(content))
			channel := Channel{ChannelID: v[0], States: v[1]}
			channels = append(channels, channel)
		}
	}

	return channels
}

func getCallCount() int {
	checkTools("bash", "asterisk", "grep")
	cmd := mapBash[211]
	// out 末尾跟了一个换行(ASCII 10)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
		return 0
	}
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)

	return count
}

/**
 * 功能: 确认有激活的信道, 统计其数量
 *
 * note: channels >= 2 * calls
 */
func getChannelCount() (int, bool) {
	cmd := mapBash[212]
	out, _ := exec.Command("bash", "-c", cmd).Output()
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)
	if count == 0 {
		return 0, false
	}
	return count, true
}

func checkTools(toolName ...string) {
	for _, tool := range toolName {
		_, err := exec.LookPath(tool)
		if err != nil {
			fmt.Println(tool, " not exist, please install it first!")
			os.Exit(1)
		}
	}
}
