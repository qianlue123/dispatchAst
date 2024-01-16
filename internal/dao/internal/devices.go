// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevicesDao is the data access object for table devices.
type DevicesDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns DevicesColumns // columns contains all the column names of Table for convenient usage.
}

// DevicesColumns defines and stores column names for table devices.
type DevicesColumns struct {
	Id           string //
	Tech         string //
	Dial         string //
	Devicetype   string //
	User         string //
	Description  string //
	EmergencyCid string //
	HintOverride string //
}

// devicesColumns holds the columns for table devices.
var devicesColumns = DevicesColumns{
	Id:           "id",
	Tech:         "tech",
	Dial:         "dial",
	Devicetype:   "devicetype",
	User:         "user",
	Description:  "description",
	EmergencyCid: "emergency_cid",
	HintOverride: "hint_override",
}

// NewDevicesDao creates and returns a new DAO object for table data access.
func NewDevicesDao() *DevicesDao {
	return &DevicesDao{
		group:   "default",
		table:   "devices",
		columns: devicesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevicesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevicesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevicesDao) Columns() DevicesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevicesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevicesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevicesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
