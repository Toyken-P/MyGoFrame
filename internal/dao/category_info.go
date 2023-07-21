// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myGoFrame/internal/dao/internal"
)

// internalCategoryInfoDao is internal type for wrapping internal DAO implements.
type internalCategoryInfoDao = *internal.CategoryInfoDao

// categoryInfoDao is the data access object for table category_info.
// You can define custom methods on it to extend its functionality as you wish.
type categoryInfoDao struct {
	internalCategoryInfoDao
}

var (
	// CategoryInfo is globally public accessible object for table category_info operations.
	CategoryInfo = categoryInfoDao{
		internal.NewCategoryInfoDao(),
	}
)

// Fill with you ideas below.
