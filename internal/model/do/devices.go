// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Devices is the golang structure of table devices for DAO operations like Where/Data.
type Devices struct {
	g.Meta       `orm:"table:devices, do:true"`
	Id           interface{} //
	Tech         interface{} //
	Dial         interface{} //
	Devicetype   interface{} //
	User         interface{} //
	Description  interface{} //
	EmergencyCid interface{} //
	HintOverride interface{} //
}
