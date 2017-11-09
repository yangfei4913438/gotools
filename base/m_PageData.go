package base

import (
	"errors"
)

// 数据分页计算器
func DataPage(all_num, page, number int64) (start_number, end_number int64, err error) {
	// 计算分页索引
	start_number = (page - 1) * number
	if start_number >= all_num {
		//起始值超过实际存在的数量，返回空的
		return 0, 0, errors.New("分页错误：起始值大于总数量! ")
	} else {
		if number > (all_num - start_number) {
			end_number = all_num
		} else {
			end_number = start_number + number
		}
	}
	// 返回结果
	return start_number, end_number, nil
}
