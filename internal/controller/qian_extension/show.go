package qian_extension

import (
	"bytes"
	Ast "dispatchAst/internal/consts"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CMD(cmd string) string {
	return fmt.Sprintf("asterisk -rx '%s' ", cmd)
}

// 核心命令 asterisk -rx 'xxx'
// NOTE 逐渐废弃
var asteriskrx = map[int]string{
	// 1-serise, core show
	1: "asterisk -rx 'core show help' ",
	2: "asterisk -rx 'core show calls' ",
	3: "asterisk -rx 'core show calls uptime' ",
	4: "asterisk -rx 'core show sysinfo' ",
	// hints 无法判断呼叫(直接当作是 inuse ) , 捕捉 ring 需要 channel 或者 pjsip
	11: "asterisk -rx 'core show hints' ",
	12: "asterisk -rx 'core show channels' ",
	13: "asterisk -rx 'core show channels count' ",

	// 2-series, database show
	21: "asterisk -rx 'database show registrar/contact' ",

	// 3-series, pjsip show
	31: "asterisk -rx 'pjsip show endpoints' ",
	32: "asterisk -rx 'pjsip show channels' ",
	// channelid 可以由
	33: "asterisk -rx 'pjsip show channel PJSIP/channelid' ",
	// 结果里, 状态是主动打电话的 ring 也包含在 in use 里
	34: "asterisk -rx 'pjsip list endpoints' ",
	// 比show多出通话时间列属性
	35: "asterisk -rx 'pjsip list channels' ",
}

// 命令二次包装, 在asterisk 命令基础上叠加 shell
// TODO 三次管道符可能得到状态码1, xargs 可能提升
var mapBash = map[int]string{
	81: CMD(Ast.RX[8]) + "| grep 'In use' | wc -l ",
	82: CMD(Ast.RX[8]) + "| grep 'In use' | awk '{print $2}' ",
	83: CMD(Ast.RX[8]) + "| grep --ignore-case 'not in use' | wc -l ",
	84: CMD(Ast.RX[8]) + "| grep --ignore-case 'not in use' | awk '{print $2}' ",
	91: CMD(Ast.RX[9]) + "| wc -l",

	571: CMD(Ast.RX[57]) + "| grep --ignore-case 'bridge ID' ",

	1100: CMD(Ast.RX[56]) + "| grep PJSIP",
	1101: CMD(Ast.RX[56]) + "| grep PJSIP | grep -i idle",
	1102: CMD(Ast.RX[56]) + "| grep -i inuse",
	1103: CMD(Ast.RX[56]) + "| grep -i ringing",
	1104: CMD(Ast.RX[56]) + "| grep -i unavailable",
	1105: CMD(Ast.RX[56]) + "| grep -i 'inuse&ringing' ", // 打给已经通话的

	1201: CMD(Ast.RX[54]) + "| grep Ring\\ ",

	2101: asteriskrx[21] + "| tail --lines 1 | awk '{print $1}'",

	// 显示 channel 的信息, 4 列
	3102: CMD(Ast.RX[8]) + "| grep Channel | tail -n -2",

	3401: CMD(Ast.RX[14]) + "| tail --lines +5 | head --lines -3 ",
	// 可用各类状态的电话数量, 先粗后细
	3402: CMD(Ast.RX[14]) + "| grep --ignore-case 'not in use' | wc -l",
	3403: CMD(Ast.RX[14]) + "| grep --ignore-case 'not in use' ",
	// 再加一次管道防 not in use 状态
	3404: CMD(Ast.RX[14]) + "| grep --ignore-case 'in use' | grep -iv 'not' | wc -l",
	3405: CMD(Ast.RX[14]) + "| grep --ignore-case 'in use' | grep -iv 'not' ",
	3406: CMD(Ast.RX[14]) + "| grep --ignore-case 'Ringing' | wc -l",
	3407: CMD(Ast.RX[14]) + "| grep --ignore-case 'Ringing' ",
	3408: CMD(Ast.RX[14]) + "| grep --ignore-case 'unavailable' | wc -l",
	3409: CMD(Ast.RX[14]) + "| grep --ignore-case 'unavailable' ",
}

// 获取电话机对象所有的信息
func GetAllExtension(state int) (arr []extension) {
	switch state {
	case Idle:
		if getCount(mapBash[3402]) > 0 {
			arr = getArrExt(mapBash[3402+1])
		}

	case InUse:
		if getCount(mapBash[3404]) > 0 {
			// 执行命令提取信息
			cmd := mapBash[3404+1] + "| awk '{print $2}'"
			out, _ := exec.Command("bash", "-c", cmd).Output()
			// 先去掉末尾换行再按换行切分
			out = out[:len(out)-1]
			outs := bytes.Split(out, []byte("\n"))
			for i, content := range outs {
				fmt.Printf("%d %v \n", i, string(content))

				var ext extension
				if strings.Contains(string(content), "/") {
					// 在格式 1234/5678 里提取 1234 当机号, 提取 5678 当 CID
					v := strings.Split(string(content), "/")
					ext = extension{ExtName: v[0], CID: v[1]}
				} else {
					ext.ExtName = string(content)
				}

				ReplenishData_ext(&ext)
				arr = append(arr, ext)
			}
		}

	case Ringing:
		if getCount(mapBash[3406]) > 0 {
			arr = getArrExt(mapBash[3406+1])
		}

	case Ring:
		//TODO 呼叫状态包含在 in use 里, 额外提取需要其他命令

	case Unavailable:
		if getCount(mapBash[3408]) > 0 {
			arr = getArrExt(mapBash[3408+1])
		}

	default:
	}

	fmt.Println(arr)
	return arr
}

// 统计电话机数量
func GetCountExt() {
	checkTools("bash", "asterisk", "fwconsole", "grep")

	cmd := mapBash[2101]
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))
}

// 34xx 系列命令专用
// asterisk -rx 'pjsip list endpoints' | grep -i ...
func getArrExt(cmd string) (arr []extension) {
	arr = make([]extension, 0)
	// 在输入命令的基础上再加 awk
	out, _ := exec.Command("bash", "-c", cmd+"| awk '{print $2}'").Output()
	// 先去掉末尾换行再按换行切分
	out = out[:len(out)-1]
	outs := bytes.Split(out, []byte("\n"))
	for i, content := range outs {
		fmt.Printf("%d %v \n", i, string(content))

		ext := extension{}
		if !strings.Contains(string(content), "/") {
			// 按照 <权威指南 5th> 示例, 创建的话机可能没有设置 CID
			ext.ExtName = string(content)
		} else {
			// 在格式 1234/5678 里提取 1234 当机号, 提取 5678 当 CID
			v := strings.Split(string(content), "/")
			ext.ExtName, ext.CID = v[0], v[1]
		}

		ReplenishData_ext(&ext)
		arr = append(arr, ext)
	}
	return arr
}

func GetExtensionState(name string) (state int) {
	state = Unavailable

	if getCount(mapBash[81]) > 0 {
		out, _ := exec.Command("bash", "-c", mapBash[82]).Output()
		out = out[:len(out)-1]
		outs := bytes.Split(out, []byte("\n"))
		for _, content := range outs {
			v := strings.Split(string(content), "/")
			if v[0] == name {
				state = InUse
			}
		}
	}

	// check array who are not in use
	if getCount(mapBash[83]) > 0 {
		out, _ := exec.Command("bash", "-c", mapBash[84]).Output()
		out = out[:len(out)-1]
		outs := bytes.Split(out, []byte("\n"))
		for _, content := range outs {
			v := strings.Split(string(content), "/")
			if v[0] == name {
				state = NotInUse
			}
		}
	}

	return state
}

func GetStateExt(state int) string {
	checkTools("bash", "asterisk", "fwconsole", "grep")
	cmd := ""
	switch state {
	case Idle:
		cmd = mapBash[1101]
	case InUse:
		cmd = mapBash[1102]
	case Ring:
		cmd = mapBash[1201]
	case Ringing:
		cmd = mapBash[1103]
	case Unavailable:
		cmd = mapBash[1104]
	default:
		cmd = mapBash[1100]
	}

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		if ins, ok := err.(*exec.ExitError); ok {
			// 命令组装异常, 返回状态码1
			fmt.Println(ins.ExitCode())
			out = []byte("0")
		} else {
			fmt.Printf("命令组装出错, 请检查: %s", err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println(string(out))
	}
	return string(out)
}

/**
 * 功能: 返回一个数
 */
func getCount(cmd string) int {
	checkTools("bash", "asterisk", "grep")

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
 * 功能: 模拟 whereis , 确保系统里有这些命令行工具
 */
func checkTools(toolName ...string) {
	for _, tool := range toolName {
		_, err := exec.LookPath(tool)
		if err != nil {
			fmt.Println(tool, " not exist, please install it first!")
			os.Exit(1)
		}
	}
}

// 功能: 确认提供的话机已经注册
func IsExistExt(extName string) bool {
	cmd := fmt.Sprintf(mapBash[91], extName)

	out, _ := exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]

	number, _ := strconv.Atoi(string(out))
	return (number != 2) // 如果有数据, 能数到150多行
}
