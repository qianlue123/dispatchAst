package qian_conference

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	Ast "dispatchAst/internal/consts"
)

// 定义枚举映射话机的状态
const (
	Idle     = iota
	NotInUse = 0
	InUse    = iota
	_
	ALL
)

func CMD(cmd string) string {
	return fmt.Sprintf("sudo asterisk -rx '%s' ", cmd)
}

// 命令二次包装, 在asterisk 命令基础上叠加 shell
var mapBash = map[int]string{
	1161: CMD(Ast.RX[116]) + "! wc -l ",           // >= 1
	1511: CMD(Ast.RX[151]) + "| wc -l ",           // >= 2
	1512: CMD(Ast.RX[151]) + "| tail --lines +3 ", // 去除字符表标题
}

func GetConfCount(state int) (count int) {
	switch state {
	case Idle:

	case InUse:
		count = GetConfCountInUse()
	default:
		count = GetConfCountAll()
	}

	return count
}

// 功能: 获取所有注册的房间号数量
func GetConfCountAll() int {
	out, _ := exec.Command("bash", "-c", mapBash[1161]).Output()
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)

	return (count - 1)
}

func GetConfCountInUse() int {
	/*
			Conference Bridge Name           Users  Marked Locked Muted
		  ================================ ====== ====== ====== =====
	*/
	out, _ := exec.Command("bash", "-c", mapBash[1511]).Output()
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)

	return (count - 2)
}

// 获取正在使用的会议房间集合
func GetRooms() []ConfBridge {
	rooms := make([]ConfBridge, 0)

	out, _ := exec.Command("bash", "-c", mapBash[1512]).Output()
	out = out[:len(out)-1]

	outs := strings.Split(string(out), "\n")
	for _, roomStr := range outs {
		// 会议房间的字符串格式内含有多个空格, 去除后只留一个
		reg := regexp.MustCompile(`\s+`)
		e := reg.Split(roomStr, -1)

		rooms = append(rooms, ConfBridge{e[0], str2int(e[1]), str2int(e[2]), e[3], e[4]})
	}

	return rooms
}

func GetEmptyRooms() []ConfBridge {
	rooms := make([]ConfBridge, 0)
	// TODO

	return rooms
}

func str2int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
