package qian_extension

import (
	"fmt"

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
