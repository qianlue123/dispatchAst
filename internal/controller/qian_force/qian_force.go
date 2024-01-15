package qian_force

import (
	"bytes"
	"context"
	"dispatchAst/api/qian_force"
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
	thing := req.GetQuery("dothing2", "0")

	// 有 extension 参数就做这件事
	if extension.String() != "0" {
		data := make([]Channel, 0)

		if count := getCallCount(); count == 0 {
			code, msg = 1, "没有电话在使用中"
		} else {
			data = GetChannelInfo(extension.String())

			if err := req.Parse(&data); err != nil {
				code, msg = 1, "error: "+err.Error()
			}
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
	channelID := req.GetQuery("channelid", "PJSIP/2024-00000066")

	cmd := fmt.Sprintf(asteriskrx[51], channelID)
	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))
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

// = callcount * 2
func getChannelCount() int {
	cmd := mapBash[212]
	out, _ := exec.Command("bash", "-c", cmd).Output()
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)
	return count
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
