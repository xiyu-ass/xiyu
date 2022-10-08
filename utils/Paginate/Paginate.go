package Paginate

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

// 分页
func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//Atoi:用于将字符串类型转换为int类型
		pageUp, _ := strconv.Atoi(page)
		if pageUp == 0 {
			pageUp = 1
		}
		pageSizeInt, _ := strconv.Atoi(pageSize)
		switch {
		case pageSizeInt <= 0:
			pageSizeInt = 10
		}
		offset := (pageUp - 1) * pageSizeInt
		return db.Offset(offset).Limit(pageSizeInt)
	}
}
