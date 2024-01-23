package qian_extension

import (
	"bytes"
	"fmt"

	SQL "dispatchAst/internal/consts"

	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
)

/**
 * 功能: 根据话机号通过查表填补话机对象的其他属性
 * 参数: ext [*Object] 直接对传入的对象补全, 不返回
 */
func ReplenishData_ext(ext *extension) {
	table1 := g.Model("devices")
	switch {
	// 优先在 devices 表中搜寻
	case ext.ExtName != "":
		device, err := table1.Where("id", ext.ExtName).
			Fields("id, user, description").
			One()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		ext.Desc = device["description"].String() // gdb.Record 转换

	default:
		//TODO
	}
}

/* ***** ***** *****  备用方案  ***** ***** ***** */

// 获取一条记录的3个字段 id,user,description
func Spare_ReplenishData_ext(ext *extension) {
	recordOne := make(map[int]string, 0)

	cmd := fmt.Sprintf(SQL.E[31]+"| sed --quiet 2p", ext.ExtName)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
		return
	}
	out = out[:len(out)-1]
	outs := bytes.Split(out, []byte("\t"))
	for i, content := range outs {
		recordOne[i] = string(content)
	}

	ext.Desc = recordOne[2]
}
