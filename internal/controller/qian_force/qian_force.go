package qian_force

import (
	"context"
	"dispatchAst/api/qian_force"
	"fmt"
	"os"
	"os/exec"
	"strconv"

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

func (c *Controller) Get(req *ghttp.Request) {
	req.Response.Writeln("asterisk channel")

	// 返回channel id, 可能有多个
	ChannelArr := GetChannelInfo()

	req.Response.WriteJson(ChannelArr)

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

func GetChannelInfo() string {
	if count := getCallCount(); count == 0 {
		return "0"
	}

	cmd := mapBash[21]
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))

	return string(out)
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
