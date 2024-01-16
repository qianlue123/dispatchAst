// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dispatchAst/internal/dao/internal"
)

// internalDevicesDao is internal type for wrapping internal DAO implements.
type internalDevicesDao = *internal.DevicesDao

// devicesDao is the data access object for table devices.
// You can define custom methods on it to extend its functionality as you wish.
type devicesDao struct {
	internalDevicesDao
}

var (
	// Devices is globally public accessible object for table devices operations.
	Devices = devicesDao{
		internal.NewDevicesDao(),
	}
)

// Fill with you ideas below.
