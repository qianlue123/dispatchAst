package qian_extension

import (
	"fmt"
	"os"
	"os/exec"
)

/*
   展示话机状态信息

*/

// 数据库命令 mysql -D -e 'xxx \G'
var mysqle = map[int]string{
	1: "mysql -e 'show databases \\G' ",
	2: "mysql -D asterisk -e 'show tables \\G' ",
	3: "mysql -D asterisk -e 'select id from devices \\G' ",

	// 反查已知列名所在的表, 可能不止一个
	11: "mysql -e \"SELECT table_name from information_schema.columns " +
		"where TABLE_SCHEMA='asterisk' and COLUMN_NAME='%s' ; \" ",
}

// 核心命令 asterisk -rx 'xxx'
var asteriskrx = map[int]string{
	// 1-serise, core show
	1: "asterisk -rx 'core show help' ",
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
	// 比show多出通话时间列属性
	34: "asterisk -rx 'pjsip list channels' ",
}

// 命令二次包装, 在asterisk 命令基础上叠加 shell
// TODO 三次管道符可能得到状态码1, xargs 可能提升
var mapBash = map[int]string{
	1100: asteriskrx[11] + "| grep PJSIP",
	1101: asteriskrx[11] + "| grep PJSIP | grep -i idle",
	1102: asteriskrx[11] + "| grep -i inuse",
	1103: asteriskrx[11] + "| grep -i ringing",
	1104: asteriskrx[11] + "| grep -i unavailable",

	1201: asteriskrx[12] + "| grep Ring\\ ",

	2101: asteriskrx[21] + "| tail --lines 1 | awk '{print $1}'",

	// 显示 channel 的信息, 4 列
	3102: asteriskrx[31] + "| grep Channel | tail -n -2",
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
