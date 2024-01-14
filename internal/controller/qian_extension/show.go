package qian_extension

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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
	1100: asteriskrx[11] + "| grep PJSIP",
	1101: asteriskrx[11] + "| grep PJSIP | grep -i idle",
	1102: asteriskrx[11] + "| grep -i inuse",
	1103: asteriskrx[11] + "| grep -i ringing",
	1104: asteriskrx[11] + "| grep -i unavailable",
	1105: asteriskrx[11] + "| grep -i 'inuse&ringing' ", // 打给已经通话的

	1201: asteriskrx[12] + "| grep Ring\\ ",

	2101: asteriskrx[21] + "| tail --lines 1 | awk '{print $1}'",

	// 显示 channel 的信息, 4 列
	3102: asteriskrx[31] + "| grep Channel | tail -n -2",

	3401: asteriskrx[34] + "| tail --lines +5 | head --lines -3 ",
	// 可用各类状态的电话数量, 先粗后细
	3402: asteriskrx[34] + "| grep --ignore-case 'not in use' | wc -l",
	3403: asteriskrx[34] + "| grep --ignore-case 'not in use' ",
	// 再加一次管道防 not in use 状态
	3404: asteriskrx[34] + "| grep --ignore-case 'in use' | grep -iv 'not' | wc -l",
	3405: asteriskrx[34] + "| grep --ignore-case 'in use' | grep -iv 'not' ",
	3406: asteriskrx[34] + "| grep --ignore-case 'Ringing' | wc -l",
	3407: asteriskrx[34] + "| grep --ignore-case 'Ringing' ",
	3408: asteriskrx[34] + "| grep --ignore-case 'unavailable' | wc -l",
	3409: asteriskrx[34] + "| grep --ignore-case 'unavailable' ",
}

// 获取电话机对象所有的信息
func GetAllExtension(state int) (arr []extension) {
	switch state {
	case Idle:
		if getCount(mapBash[3402]) > 0 {
			// 执行命令提取信息
			cmd := mapBash[3402+1]
			out, _ := exec.Command("bash", "-c", cmd).Output()
			out = out[:len(out)-1]

			fmt.Println(string(out))
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
				//TODO 在格式 1234/5678 里提取 1234
			}
		}
	default:
	}
	return []extension{}
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
