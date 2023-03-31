package utils

import "math"

func TotalPage(total int64, pageSize int64) int64 {
	return int64(math.Ceil(float64(total) / float64(pageSize)))
}
